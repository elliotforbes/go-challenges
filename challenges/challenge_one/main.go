package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"text/tabwriter"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

/*
Utilities.
*/

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

/*
Functionality.
*/

func printCPUUtilization() {
	fmt.Println("\n\t\t CPU Utilization")

	v, err := mem.VirtualMemory()
	check(err)

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "Total", "Free", "Used Percent")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t", "----", "----", "----")
	fmt.Fprintf(w, "\n %v\t%v\t%f\n", v.Total, v.Free, v.UsedPercent)
}

func printRAMUtilization() {
	fmt.Println("\n\t\t RAM Utilization")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Alloc", "TotalAlloc", "Sys", "NumGC")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "----", "----", "----", "----")
	fmt.Fprintf(w, "\n %v%s\t%v%s\t%v%s\t%v\n",
		bToMb(m.Alloc), " MB",
		bToMb(m.TotalAlloc), " MB",
		bToMb(m.Sys), " MB",
		m.NumGC)
}

func printBackingStoreUtilization() {
	fmt.Println("\n\t\t Backing Store Utilization")

	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)

		usage = append(usage, u)
		printUsage(u)
	}
}

func printUsage(u *disk.UsageStat) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	if len(u.Path) > 18 {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Drive/Path", "Percentage Full", "Total", "Free", "Used")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\n",
			u.Path,
			strconv.FormatFloat(u.UsedPercent, 'f', 2, 64)+"%",
			strconv.FormatUint(u.Total/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Free/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Used/1024/1024/1024, 10)+" GB")
	} else {
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\t", "Drive/Path", "Percentage Full", "Total", "Free", "Used")
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")
		fmt.Fprintf(w, "\n %s\t\t%s\t%s\t%s\t%s\n",
			u.Path,
			strconv.FormatFloat(u.UsedPercent, 'f', 2, 64)+"%",
			strconv.FormatUint(u.Total/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Free/1024/1024/1024, 10)+" GB",
			strconv.FormatUint(u.Used/1024/1024/1024, 10)+" GB")
	}
}

func main() {
	printCPUUtilization()
	printRAMUtilization()
	printBackingStoreUtilization()
}
