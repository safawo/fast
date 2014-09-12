package run

import (
	"fmt"
	"time"
)

type fastRunInfo struct {
	startTime string
	rootPath  string
}

var runInfo fastRunInfo = fastRunInfo{}

func init() {
	fmt.Println("Start Fast Run")

	runInfo.startTime = time.Now().Format("2006-01-02 15:04:05")

	fmt.Println("  Start at ", runInfo.startTime)
}

func StartTime() (startTime string) {
	return runInfo.startTime
}

func RootPath() (rootPath string) {
	return runInfo.rootPath
}
