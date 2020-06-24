package main

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
)

var KB = uint64(1024)

// DiskThreshold is the warning threshold for disks
const DiskThreshold = 90

func checkDiskSpace(disk string) (bool, string, string, error) {
	usage := du.NewDiskUsage(disk)

	msg := fmt.Sprintf("Available: %dkb; Total: %dkb, Usage: %.2f%%", usage.Available()/(KB*KB), usage.Size()/(KB*KB), usage.Usage()*100)

	okay := true

	// send false if less than 10% disk space left
	if usage.Usage()*100 > DiskThreshold {
		okay = false
	}

	data := fmt.Sprintf("%d|%d", usage.Available()/(KB*KB), usage.Size()/(KB*KB))
	return okay, msg, data, nil
}
