package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
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

func memoryUsage() (string, string) {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	totalMemory := bytefmt.ByteSize(memory.Total)
	percentageMemory := string(memory.Used * 100 / memory.Total)

	return totalMemory, percentageMemory
}

func cpuUsage() string {
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

	percentageCpu := strconv.FormatFloat(100-float64(after.Idle-before.Idle)/total*100, 'E', -1, 64)

	return percentageCpu
}

func loadAvarage() (string, string, string) {
	load, err := loadavg.Get()
	checkError(err)

	loadAverage1 := strconv.FormatFloat(float64(load.Loadavg1), 'E', -1, 64)
	loadAverage5 := strconv.FormatFloat(float64(load.Loadavg5), 'E', -1, 64)
	loadAverage15 := strconv.FormatFloat(float64(load.Loadavg15), 'E', -1, 64)

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

func markdownGenerator(uptime time.Duration, percentagememory string, percentagecpu string,
	loadaverage1 string, loadaverage5 string, loadaverage15 string, percentagedisk string) {
	basicTable, _ := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build("Up Time", "Memory Usage Percentage", "CPU Usage Percentage", "LoadAVG1", "LoadAVG5", "LoadAVG15", "Disk Usage Percentage").
		Format([][]string{
			{string(uptime), percentagememory, percentagecpu, loadaverage1, loadaverage5, loadaverage15, percentagedisk},
		})
	fmt.Print(basicTable)
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

	markdownGenerator(uptime, percentagememory, percentagecpu, loadaverage1, loadaverage5, loadaverage15, percentagedisk)
}
