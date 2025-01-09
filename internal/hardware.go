package internal

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Hostname  string `json:"hostname"`
	CPU       string `json:"cpu"`
	CPUCores  string `json:"cpucores"`
	Platform  string `json:"platform"`
	RAMTotal  string `json:"ramtotal"`
	RAMFree   string `json:"ramfree"`
	DiskTotal string `json:"disktotal"`
	DiskFree  string `json:"diskfree"`
}

func GetCPUSection() []cpu.InfoStat {
	cpuStat, err := cpu.Info()
	if err != nil {
		return nil
	}
	return cpuStat
}

func GetDiskSection() *disk.UsageStat {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil
	}
	return diskStat
}

func GetHostSection() *host.InfoStat {
	hostStat, err := host.Info()
	if err != nil {
		return nil
	}
	return hostStat
}

func GetMemSection() *mem.VirtualMemoryStat {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil
	}
	return vmStat
}

func GetSystemInfo() *SystemInfo {
	vmStat := GetMemSection()
	hostStat := GetHostSection()
	cpuStat := GetCPUSection()
	diskStat := GetDiskSection()

	sysInfo := &SystemInfo{
		hostStat.Hostname,
		cpuStat[0].ModelName,
		fmt.Sprintf("%d", len(cpuStat)),
		hostStat.Platform,
		fmt.Sprintf("%d", vmStat.Total),
		fmt.Sprintf("%d", vmStat.Free),
		fmt.Sprintf("%d", diskStat.Total),
		fmt.Sprintf("%d", diskStat.Free),
	}

	return sysInfo
}
