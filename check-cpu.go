package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

const CPUWarningThreshold = 50.0
const CPUThreshold = 85.0

func checkCPU() (bool, string, string, int, error) {

	c, _ := cpu.Percent(0, false)

	okay := true
	newStatusID := StatusHealthy
	msg := fmt.Sprintf("CPU usage okay: %0.4f%%", c[0])

	if c[0] > CPUWarningThreshold {
		newStatusID = StatusWarning
		msg = fmt.Sprintf("Warning: Moderate CPU usage: %0.4f%%", c[0])
		okay = false
	} else if c[0] > CPUThreshold {
		okay = false
		msg = fmt.Sprintf("Problem: High CPU usage %0.4f%%", c[0])
		newStatusID = StatusProblem
	}

	return okay, msg, "", newStatusID, nil
}
