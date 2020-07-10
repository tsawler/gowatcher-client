package main

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
)

// DiskThreshold is the threshold for disk space
const DiskThreshold = 90

// DiskWarningThreshold is the warning threshold for disk space
const DiskWarningThreshold = 75

// MB is a megabyte
var MB = uint64(1024)

// checkDiskSpace performs the disk space check
func checkDiskSpace(disk string) (bool, string, string, int, error) {
	usage := du.NewDiskUsage(disk)

	msg := fmt.Sprintf("%s: Available: %d MB; Total: %d MB, Usage: %.2f%%", disk, usage.Available()/(MB*MB), usage.Size()/(MB*MB), usage.Usage()*100)

	okay := true
	newStatusID := 1

	// send false if less than 10% disk space left
	if usage.Usage()*100 > DiskThreshold {
		// problem
		okay = false
		newStatusID = 2
	} else if usage.Usage()*100 > DiskWarningThreshold {
		// warning
		okay = false
		newStatusID = 4
	}

	data := fmt.Sprintf("%d|%d", usage.Available()/(MB*MB), usage.Size()/(MB*MB))
	return okay, msg, data, newStatusID, nil
}
