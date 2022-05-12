package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

func memoryUsage() (string, string) {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	totalMemory := bytefmt.ByteSize(memory.Total)
	percentageMemory := strconv.FormatUint(memory.Used*100/memory.Total, 10)

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

	percentageCpu := fmt.Sprintf("%.2f", 100-float64(after.Idle-before.Idle)/total*100)

	return percentageCpu
}

func loadAvarage() (string, string, string) {
	load, err := loadavg.Get()
	checkError(err)

	loadAverage1 := fmt.Sprintf("%.2f", float64(load.Loadavg1))
	loadAverage5 := fmt.Sprintf("%.2f", float64(load.Loadavg5))
	loadAverage15 := fmt.Sprintf("%.2f", float64(load.Loadavg15))

	return loadAverage1, loadAverage5, loadAverage15
}

func upTime() time.Duration {
	uptime, err := uptime.Get()
	checkError(err)

	return uptime
}

func diskUsage() (string, string) {
	usage := du.NewDiskUsage(".")
	diskSize := bytefmt.ByteSize(usage.Size())
	percentageDisk := fmt.Sprintf("%.2f", usage.Usage()*100)

	return diskSize, percentageDisk
}

func markdownGenerator(hostname string, uptime string, percentagecpu string, percentagedisk string,
	percentagememory string, loadavarage1 string, loadaverage5 string, loadaverage15 string) {
	basicTable, _ := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build("Hostname", "Up Time", "CPU Usage Percentage", "Disk Usage Percentage",
			"Memory Usage Percentage", "Load Average 1", "Load Average 5", "Load average 15").
		Format([][]string{
			{hostname, uptime, percentagecpu, percentagedisk, percentagememory, loadavarage1, loadaverage5, loadaverage15},
		})
	fmt.Print(basicTable)
}

func getIp() *net.UDPAddr {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)
	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	return ipAddress
}

func main() {
	totalmemory, percentagememory := memoryUsage()
	percentagecpu := cpuUsage()
	loadaverage1, loadaverage5, loadaverage15 := loadAvarage()
	uptime := upTime().String()
	disksize, percentagedisk := diskUsage()
	hostname, _ := os.Hostname()

	fmt.Printf("\nSystem uptime is: %v\n", uptime)
	fmt.Printf("\nHostname is: %v\n", hostname)
	fmt.Printf("\nMemory usage percentage is %v %% out of %v GB\n", percentagememory, totalmemory)
	fmt.Printf("\nCPU usage percentage is %v %%\n", percentagecpu)
	fmt.Printf("\nLoad Avagerate for last minute is %v\n", loadaverage1)
	fmt.Printf("Load Avagerate for last 5 minutes is %v\n", loadaverage5)
	fmt.Printf("Load Avagerate for last 15 minutes is %v\n", loadaverage15)
	fmt.Printf("\nDisk usage percentage is %v %% out of %v GB\n\n", percentagedisk, disksize)

	markdownGenerator(hostname, uptime, percentagecpu, percentagedisk, percentagememory, loadaverage1, loadaverage5, loadaverage15)

}
