package get

import (
	"commander/error"
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNameSpaces(k8s *kubernetes.Clientset) []string {
	nameSpaces := []string{}
	nameSpaceList, err := k8s.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	error.CheckErr(err)

	for _, nameSpace := range nameSpaceList.Items {
		nameSpaces = append(nameSpaces, nameSpace.Name)
	}
	fmt.Println("********from fucntion**********")
	fmt.Println(nameSpaces)
	fmt.Println("********from fucntion**********")
	return nameSpaces
}

func GetPod(k8s *kubernetes.Clientset, nameSpace string, appName string) {
	podList, err := k8s.CoreV1().Pods(nameSpace).List(context.Background(), metav1.ListOptions{})
	error.CheckErr(err)
	for _, podName := range podList.Items {
		fmt.Println(podName.Name)
	}
}

func UserInput() (string, string) {
	var nameSpace string
	var appName string

	fmt.Println("Enter the target namespace:")
	fmt.Scan(&nameSpace)
	fmt.Println("Enter the application name:")
	fmt.Scan(&appName)

	return nameSpace, appName
}
