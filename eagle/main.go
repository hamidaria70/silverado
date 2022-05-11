package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/uptime"
	"github.com/ricochet2200/go-disk-usage/du"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sshConnect() {
	cmd := exec.Command("ssh", "office", "echo 'blah...blah....blah' > golang.txt")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	checkError(err)

}

func readSshConfig() {
	home, _ := os.UserHomeDir()
	sshConfig := (filepath.Join(home, ".ssh", "config"))
	file, err := ioutil.ReadFile(sshConfig)

	checkError(err)
	fmt.Println(string(file))
}

func memoryUsage() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Printf("memory total: %v\n", bytefmt.ByteSize(memory.Total))
	fmt.Printf("memory used: %v bytes\n", bytefmt.ByteSize(memory.Used))
	fmt.Printf("memory cached: %v bytes\n", bytefmt.ByteSize(memory.Cached))
	fmt.Printf("memory free: %v bytes\n", bytefmt.ByteSize(memory.Free))
}

func cpuUsage() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)

	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}

func loadAvarage() {
	load1, err := loadavg.Get()
	checkError(err)

	fmt.Printf("load1 : %f\n", float64(load1.Loadavg1))
	fmt.Printf("load5 : %f\n", float64(load1.Loadavg5))
	fmt.Printf("load15 : %f\n", float64(load1.Loadavg15))
}

func upTime() {
	uptime, err := uptime.Get()
	checkError(err)

	fmt.Println(uptime)
}

func diskUsage() {
	usage := du.NewDiskUsage(".")
	fmt.Println("Free:", bytefmt.ByteSize(usage.Free()))
	fmt.Println("Available:", bytefmt.ByteSize(usage.Available()))
	fmt.Println("Size:", bytefmt.ByteSize(usage.Size()))
	fmt.Println("Used:", bytefmt.ByteSize(usage.Used()))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}

func main() {
	readSshConfig()
	sshConnect()
	memoryUsage()
	cpuUsage()
	loadAvarage()
	upTime()
	diskUsage()
}
