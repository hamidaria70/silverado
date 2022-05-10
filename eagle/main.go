package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	home, _ := os.UserHomeDir()
	sshConfig := (filepath.Join(home, ".ssh", "config"))
	file, err := ioutil.ReadFile(sshConfig)
	
	checkError(err)
	fmt.Println(string(file))
}
