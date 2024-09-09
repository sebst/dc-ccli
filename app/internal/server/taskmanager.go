package server

import (
	"time"

	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	PID         int32
	Name        string
	CPUPercent  float64
	MemoryUsage uint64
	User        string
	StartTime   time.Time
}

func getRunningProcesses() ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var processList []ProcessInfo
	for _, proc := range procs {
		pid := proc.Pid

		// Get process name
		name, err := proc.Name()
		if err != nil {
			name = "unknown"
		}

		// Get CPU usage percentage
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		// Get memory usage in bytes
		memInfo, err := proc.MemoryInfo()
		memoryUsage := uint64(0)
		if err == nil && memInfo != nil {
			memoryUsage = memInfo.RSS
		}

		// Get user who started the process
		user, err := proc.Username()
		if err != nil {
			user = "unknown"
		}

		// Get process start time
		startTime := time.Time{}
		createTime, err := proc.CreateTime()
		if err == nil {
			startTime = time.Unix(0, createTime*int64(time.Millisecond))
		}

		processList = append(processList, ProcessInfo{
			PID:         pid,
			Name:        name,
			CPUPercent:  cpuPercent,
			MemoryUsage: memoryUsage,
			User:        user,
			StartTime:   startTime,
		})
	}

	return processList, nil
}
