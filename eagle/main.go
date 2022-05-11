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

func memoryUsage() (string, uint64) {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	totalMemory := bytefmt.ByteSize(memory.Total)
	percentageMemory := memory.Used * 100 / memory.Total

	return totalMemory, percentageMemory
}

func cpuUsage() float64 {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	total := float64(after.Total - before.Total)

	percentageCpu := 100 - float64(after.Idle-before.Idle)/total*100

	return percentageCpu
}

func loadAvarage() (float64, float64, float64) {
	load, err := loadavg.Get()
	checkError(err)

	loadAverage1 := float64(load.Loadavg1)
	loadAverage5 := float64(load.Loadavg5)
	loadAverage15 := float64(load.Loadavg15)

	return loadAverage1, loadAverage5, loadAverage15
}

func upTime() time.Duration {
	uptime, err := uptime.Get()
	checkError(err)

	fmt.Printf("System uptime is %v: \n", uptime)
	return uptime
}

func diskUsage() (string, string) {
	usage := du.NewDiskUsage(".")
	diskSize := bytefmt.ByteSize(usage.Size())
	percentageDisk := fmt.Sprintf("%v", usage.Usage()*100)

	return diskSize, percentageDisk
}

func main() {
	readSshConfig()
	sshConnect()
	totalmemory, percentagememory := memoryUsage()
	percentagecpu := cpuUsage()
	loadaverage1, loadaverage5, loadaverage15 := loadAvarage()
	uptime := upTime()
	disksize, percentagedisk := diskUsage()

	fmt.Printf("\nSystem uptime is: %v\n", uptime)
	fmt.Printf("\nMemory usage percentage is %v %% out of %v GB\n", percentagememory, totalmemory)
	fmt.Printf("\nCPU usage percentage is %v %%\n", percentagecpu)
	fmt.Printf("\nLoad Avagerate for last minute is %v\n", loadaverage1)
	fmt.Printf("Load Avagerate for last 5 minutes is %v\n", loadaverage5)
	fmt.Printf("Load Avagerate for last 15 minutes is %v\n", loadaverage15)
	fmt.Printf("\nDisk usage percentage is %v %% out of %v GB\n", percentagedisk, disksize)
}
