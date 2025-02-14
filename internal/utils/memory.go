package utils

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/mem"
)

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
)

func PrintMemstats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	vm, _ := mem.VirtualMemory()
	fmt.Printf("heap in use: %d MB\n", memStats.HeapInuse/MB)
	fmt.Printf("GC nums: %d\n", memStats.NumGC)
	fmt.Printf("vm used: %d MB\n", vm.Used/MB)
}
