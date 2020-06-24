package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
)

// checkMemory checks available memory for OS
func checkMemory() (bool, string, string, error) {
	okay := false

	mem := sigar.Mem{}

	err := mem.Get()

	if err != nil {
		return false, err.Error(), "", err
	}

	usage := mem.Used / mem.Total
	if usage < MemoryThreshold {
		okay = true
	}

	return okay,
		fmt.Sprintf("Total: %d; Used: %d; Free: %d",
			format(mem.Total),
			format(mem.Used),
			format(mem.Free)),
		fmt.Sprintf("%d|%d|%d",
			format(mem.Total),
			format(mem.Used),
			format(mem.Free)), nil

}

// format just formats things nicely
func format(val uint64) uint64 {
	return val / 1024
}
