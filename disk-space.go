package main

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
)

// KB is a kilobyte
var MB = uint64(1024)

// checkDiskSpace peforms the disk space check
func checkDiskSpace(disk string) (bool, string, string, error) {
	usage := du.NewDiskUsage(disk)

	msg := fmt.Sprintf("%s: Available: %d MB; Total: %d MB, Usage: %.2f%%", disk, usage.Available()/(MB*MB), usage.Size()/(MB*MB), usage.Usage()*100)

	okay := true

	// send false if less than 10% disk space left
	if usage.Usage()*100 > DiskThreshold {
		okay = false
	}

	data := fmt.Sprintf("%d|%d", usage.Available()/(MB*MB), usage.Size()/(MB*MB))
	return okay, msg, data, nil
}
