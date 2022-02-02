package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	kubeConfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)

    data, err := ioutil.ReadFile(kubeConfig)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}