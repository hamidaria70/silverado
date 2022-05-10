package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

	cmd := exec.Command("ssh", "office", "echo 'blah...blah....blah' > golang.txt")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
