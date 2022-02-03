package main

import (
	"commander/config"
	"context"
	"fmt"
	"log"
	"path/filepath"

	//	"k8s.io/client-go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeConfigFile := config.GetConfig()
	fmt.Printf("Path to the kubeConfig is %v.\n", kubeConfigFile)
	kubeConfigData := config.ShowConfig(kubeConfigFile)
	fmt.Println(kubeConfigData)
	kubeConfig, err := clientcmd.BuildConfigFromFlags(
		"", filepath.Join(homedir.HomeDir(), ".kube", "config"),
	)
	checkErr(err)

	k8s, err := kubernetes.NewForConfig(kubeConfig)
	checkErr(err)

	nodeList, err := k8s.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	checkErr(err)

	for _, nameSpace := range nodeList.Items {
		fmt.Println(nameSpace.Name)
	}
	//	podList , err := clientset.CoreV1().pods("").List(context.Background(),metav1.ListOptions{})
	//	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}