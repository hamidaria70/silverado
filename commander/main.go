package main

import (
	"os"
	"fmt"
	"io/ioutil"
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