package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
)

func checkMemory() (bool, string, string, error) {
	okay := false

	mem := sigar.Mem{}

	err := mem.Get()

	if err != nil {
		return false, err.Error(), "", err
	}

	usage := mem.Used / mem.Total
	if usage < 80 {
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

func format(val uint64) uint64 {
	return val / 1024
}
