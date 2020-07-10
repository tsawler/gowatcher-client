package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

const CPUWarningThreshold = 50.0
const CPUThreshold = 85.0

func checkCPU() (bool, string, string, int, error) {

	c, err := cpu.Percent(0, false)
	if err != nil {
		return false, fmt.Sprintf("Error checking CPU: %s", err.Error()), "", StatusProblem, err
	}

	numCPU, err := cpu.Counts(true)
	if err != nil {
		return false, fmt.Sprintf("Error checking CPU: %s", err.Error()), "", StatusProblem, err
	}

	okay := true
	newStatusID := StatusHealthy

	msg := fmt.Sprintf("CPU usage okay: %0.4f%% average usage for %d cpu(s)", c[0], numCPU)

	if c[0] > CPUWarningThreshold {
		newStatusID = StatusWarning
		msg = fmt.Sprintf("Warning: Moderate CPU usage: %0.4f%% for %d cpu(s)", c[0], numCPU)
		okay = false
	} else if c[0] > CPUThreshold {
		okay = false
		msg = fmt.Sprintf("Problem: High CPU usage %0.4f%% for %d cpu(s)", c[0], numCPU)
		newStatusID = StatusProblem
	}

	return okay, msg, "", newStatusID, nil
}
