package main

import (
	"os"

	"github.com/google/martian/v3/log"
	"github.com/unconditionalday/source-checker/cmd"
)

var (
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = "unknown"
	osArch    = "unknown"
)

func main() {
	versions := map[string]string{
		"version":   version,
		"gitCommit": gitCommit,
		"buildTime": buildTime,
		"goVersion": goVersion,
		"osArch":    osArch,
	}

	if err := cmd.NewRootCommand(versions).Execute(); err != nil {
		log.Errorf("Source Checker failed")
		os.Exit(1)
	}
}
