package get

import (
	"commander/error"
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNode(k8s *kubernetes.Clientset) {
	nodeList, err := k8s.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	error.CheckErr(err)

	for _, nameSpace := range nodeList.Items {
		fmt.Println(nameSpace.Name)
	}
}

func GetPod(k8s *kubernetes.Clientset) {
	podList, err := k8s.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	error.CheckErr(err)
	for _, podName := range podList.Items {
		fmt.Println(podName)
	}
}
