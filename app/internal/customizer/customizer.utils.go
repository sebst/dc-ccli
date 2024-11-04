/*
Copyright © 2024 devcontainer.com
*/
package customizer

import "os"

func getBasePath() string {
	env_home := os.Getenv("HOME")
	env_dc_base := os.Getenv("DC_BASE")
	if env_dc_base != "" {
		return env_dc_base
	}
	return env_home
}
