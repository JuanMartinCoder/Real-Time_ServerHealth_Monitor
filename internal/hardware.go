package internal

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Hostname       string `json:"hostname"`
	CPU            string `json:"cpu"`
	CPUCores       string `json:"cpucores"`
	CPUSpeed       string `json:"cpuspeed"`
	Platform       string `json:"platform"`
	RAMTotal       string `json:"ramtotal"`
	RAMFree        string `json:"ramfree"`
	RAMUsedPercent string `json:"ramused"`
	DiskTotal      string `json:"disktotal"`
	DiskFree       string `json:"diskfree"`
	DiskUsed       string `json:"diskused"`
	DiskPercent    string `json:"diskpercent"`
}

const (
	megabyteDiv uint64 = 1024 * 1024
	gigabyteDiv uint64 = megabyteDiv * 1024
)

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

	RAMTotalUnitFormated := strconv.FormatUint(vmStat.Total/megabyteDiv, 10)
	RAMFreeUnitFormated := strconv.FormatUint(vmStat.Free/megabyteDiv, 10)

	DiskTotalUnitFormated := strconv.FormatUint(diskStat.Total/gigabyteDiv, 10)
	DiskFreeUnitFormated := strconv.FormatUint(diskStat.Free/gigabyteDiv, 10)
	DiskUsedUnitFormated := strconv.FormatUint(diskStat.Used/gigabyteDiv, 10)

	sysInfo := &SystemInfo{
		hostStat.Hostname,
		cpuStat[0].ModelName,
		fmt.Sprintf("%d", len(cpuStat)),
		fmt.Sprintf("%.2f", cpuStat[0].Mhz),
		hostStat.Platform,
		RAMTotalUnitFormated,
		RAMFreeUnitFormated,
		fmt.Sprintf("%.2f", vmStat.UsedPercent),
		DiskTotalUnitFormated,
		DiskFreeUnitFormated,
		DiskUsedUnitFormated,
		fmt.Sprintf("%.2f", diskStat.UsedPercent),
	}

	return sysInfo
}
