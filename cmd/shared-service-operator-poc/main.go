package main

import (
	"context"
	"runtime"

	"github.com/aerogear/shared-service-operator-poc/pkg/shared"
	clientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/operator-framework/operator-sdk/pkg/k8sclient"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	"github.com/sirupsen/logrus"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
	logrus.Infof("clientset: %+v", clientset.clientset)
}

func main() {
	printVersion()

	resource := "aerogear.org/v1alpha1"
	SharedServicekind := "SharedService"
	SharedServiceSlicekind := "SharedServiceSlice"
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		logrus.Fatalf("Failed to get watch namespace: %v", err)
	}
	resyncPeriod := 5
	logrus.Infof("Watching %s, %s, %s, %d", resource, SharedServicekind, namespace, resyncPeriod)
	sdk.Watch(resource, SharedServicekind, namespace, resyncPeriod)
	sdk.Watch(resource, SharedServiceSlicekind, namespace, resyncPeriod)
	k8client := k8sclient.GetKubeClient()
	resourceClient, _, err := k8sclient.GetResourceClient(resource, SharedServicekind, namespace)
	sdk.Handle(shared.NewHandler(k8client, resourceClient, "default"))
	sdk.Run(context.TODO())
}