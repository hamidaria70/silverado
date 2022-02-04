package main

import (
	"commander/error"
	"commander/get"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(
		"", filepath.Join(homedir.HomeDir(), ".kube", "config"),
	)
	error.CheckErr(err)

	k8s, err := kubernetes.NewForConfig(kubeConfig)
	error.CheckErr(err)

	nameSpace, appName := get.UserInput()
	get.GetNode(k8s)

	get.GetPod(k8s, nameSpace, appName)

}
