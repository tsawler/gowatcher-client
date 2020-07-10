package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

// LoadWarningThreshold is the warning threshold
const LoadWarningThreshold = 1.5

// LoadThreshold is the problem threshold
const LoadThreshold = 2.0

func checkLoad() (bool, string, string, int, error) {
	// get the system load
	systemLoad, err := load.Avg()
	if err != nil {
		return false, fmt.Sprintf("Error system load: %s", err.Error()), "", StatusProblem, err
	}

	// count the cpus
	numCPU, err := cpu.Counts(true)
	if err != nil {
		return false, fmt.Sprintf("Error checking CPU: %s", err.Error()), "", StatusProblem, err
	}

	okay := true
	newStatusID := StatusHealthy

	msg := fmt.Sprintf("System load okay: %0.4f, %0.4f, %0.4f for %d cpus", systemLoad.Load1, systemLoad.Load5, systemLoad.Load15, numCPU)

	// check the threshold * number of cpus to determine if there is a warning/problem
	if systemLoad.Load1 > LoadWarningThreshold*float64(numCPU) {
		newStatusID = StatusWarning
		msg = fmt.Sprintf("Warning: Moderate system load: %0.4f, %0.4f, %0.4f for %d cpus", systemLoad.Load1, systemLoad.Load5, systemLoad.Load15, numCPU)
		okay = false
	} else if systemLoad.Load1 > LoadThreshold*float64(numCPU) {
		okay = false
		msg = fmt.Sprintf("Problem: High system load: %0.4f, %0.4f, %0.4f for %d cpus", systemLoad.Load1, systemLoad.Load5, systemLoad.Load15, numCPU)
		newStatusID = StatusProblem
	}

	return okay, msg, "", newStatusID, nil
}
