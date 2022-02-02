package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetConfig() string {
	kubeConfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	return kubeConfig
}

func ShowConfig(kubeConfig string) string {
	data, err := ioutil.ReadFile(kubeConfig)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}
