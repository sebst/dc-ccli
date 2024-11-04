/*
Copyright © 2024 devcontainer.com
*/
package customizer

import (
	"fmt"
)

func ApplyEnv(config *Config) error {
	fmt.Println("Applying environment variables")
	for _, env := range config.Environment {
		fmt.Println("Setting environment variable:", env.Name, env.Value)
		fmt.Println("TODO: This does not work, yet") // TODO
		// // if err := os.Setenv(env.Name, env.Value); err != nil {
		// // 	return fmt.Errorf("failed to set environment variable: %w", err)
		// // }
		// command := "eval \"export " + env.Name + "=" + env.Value + "\""
		// cmd := exec.Command("bash", "-c", command)
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		// err := cmd.Run()
		// if err != nil {
		// 	return fmt.Errorf("failed to run command: %w", err)
		// }

	}
	return nil
}
