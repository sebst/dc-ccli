/*
Copyright © 2024 devcontainer.com
*/
package main

import (
	"devcontainer.com/ccli/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
