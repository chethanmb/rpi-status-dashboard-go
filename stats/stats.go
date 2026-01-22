package stats

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type Status struct {
	CPUUsage float64 `json:"cpu"`
	Temp     float64 `json:"temp"`
	MemUsed  uint64  `json:"mem_used"`
	MemTotal uint64  `json:"mem_total"`
	DiskUsed uint64  `json:"disk_used"`
	DiskTotal uint64 `json:"disk_total"`
	Uptime   uint64  `json:"uptime"`
	Hostname string  `json:"hostname"`
	Time     string  `json:"time"`
}

func Get() (*Status, error) {
	cpuP, _ := cpu.Percent(0, false)
	vm, _ := mem.VirtualMemory()
	ds, _ := disk.Usage("/")
	h, _ := host.Info()

	return &Status{
		CPUUsage: cpuP[0],
		Temp:     readTemp(),
		MemUsed:  vm.Used,
		MemTotal: vm.Total,
		DiskUsed: ds.Used,
		DiskTotal: ds.Total,
		Uptime:   h.Uptime,
		Hostname: h.Hostname,
		Time:     time.Now().Format("15:04:05"),
	}, nil
}

func readTemp() float64 {
	b, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0
	}
	val := 0.0
	fmt.Sscanf(strings.TrimSpace(string(b)), "%f", &val)
	return val / 1000
}
