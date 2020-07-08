package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
)

// checkMemory checks available memory for OS
func checkMemory() (bool, string, string, int, error) {
	okay := false

	mem := sigar.Mem{}
	newStatusID := 0
	err := mem.Get()

	if err != nil {
		return false, err.Error(), "", 2, err
	}

	usage := mem.Used / mem.Total
	if usage < MemoryThreshold {
		okay = true
		newStatusID = 1
	} else if usage < MemoryWarningThreshold {
		okay = true
		newStatusID = 4
	}

	return okay,
		fmt.Sprintf("Total: %d; Used: %d; Free: %d",
			format(mem.Total),
			format(mem.Used),
			format(mem.Free)),
		fmt.Sprintf("%d|%d|%d",
			format(mem.Total),
			format(mem.Used),
			format(mem.Free)), newStatusID, nil

}

// format just formats things nicely
func format(val uint64) uint64 {
	return val / 1024
}
