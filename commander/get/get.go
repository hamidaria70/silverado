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
	return nameSpaces
}

func GetPod(k8s *kubernetes.Clientset, nameSpace string, appName string, nameSpaces []string) {
	podList, err := k8s.CoreV1().Pods(nameSpace).List(context.Background(), metav1.ListOptions{})
	error.CheckErr(err)
	for _, podName := range podList.Items {
		fmt.Println(podName.Name)
	}
}

func UserInput(nameSpaces []string) (string, string) {
	var nameSpace string
	var appName string

	fmt.Print("Enter the target namespace: ")
	fmt.Scan(&nameSpace)
	NameSpaceChecker(nameSpace, nameSpaces)
	fmt.Print("Enter the application name: ")
	fmt.Scan(&appName)

	return nameSpace, appName
}

func NameSpaceChecker(nameSpace string, nameSpaces []string) {
	booler := false
	for _, name := range nameSpaces {
		if nameSpace == name {
			booler = true
			break
		}
	}
	error.NameSpaceCheckErr(nameSpace, booler)
}
