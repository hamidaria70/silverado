package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	kubeConfig := getConfig()
	fmt.Printf("Path to the kubeConfig is %v.\n", kubeConfig)
	kubeConfigData := showConfig(kubeConfig)
	fmt.Println(kubeConfigData)

}

func getConfig() string {
	kubeConfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	return kubeConfig
}

func showConfig(kubeConfig string) string {
    data, err := ioutil.ReadFile(kubeConfig)

    if err != nil {
        fmt.Println("File reading error", err)
    }
    return string(data)
}