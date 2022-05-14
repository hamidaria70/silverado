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

func markdownGenerator(hostname string, ip, uptime string, percentagecpu string, percentagedisk string,
	disksize string, totalmemory string, percentagememory string,
	loadavarage1 string, loadaverage5 string, loadaverage15 string) {
	basicTable, _ := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build("Hostname", "IP Address", "Up Time", "CPU Usage Percentage", "Disk Usage Percentage",
			"Memory Usage Percentage", "Load Average 1", "Load Average 5", "Load average 15").
		Format([][]string{
			{hostname, ip, uptime, percentagecpu, percentagedisk, percentagememory, loadavarage1, loadaverage5, loadaverage15},
		})

	f, err := os.Create(fmt.Sprintf("%v-data.md", ip))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(basicTable)

	if err2 != nil {
		log.Fatal(err2)
	}

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

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	totalmemory, percentagememory := memoryUsage()
	percentagecpu := cpuUsage()
	loadaverage1, loadaverage5, loadaverage15 := loadAvarage()
	uptime := upTime().String()
	disksize, percentagedisk := diskUsage()
	hostname, _ := os.Hostname()
	ip := GetOutboundIP().String()

	markdownGenerator(hostname, ip, uptime, percentagecpu, percentagedisk, disksize,
		percentagememory, totalmemory, loadaverage1, loadaverage5, loadaverage15)
}
