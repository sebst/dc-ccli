/*
Copyright © 2024 devcontainer.com
*/
package customizer

import "os"

func ApplySshKeys(config *Config) error {
	// Loop through each SSH key in the config
	for _, sshKey := range config.SshKeys {
		// Write the private key to the user's .ssh directory

		privateKeyFile := File{
			Content: Content{Text: sshKey.PrivateKey},
			Path:    ".ssh/id_rsa",
		}
		if err := createFile(privateKeyFile.Path, &privateKeyFile); err != nil {
			return err
		}
		// set the permissions of the private key file
		if err := os.Chmod(privateKeyFile.Path, 0600); err != nil {
			return err
		}

		publicKeyFile := File{
			Content: Content{Text: sshKey.PublicKey},
			Path:    ".ssh/id_rsa.pub",
		}
		if err := createFile(publicKeyFile.Path, &publicKeyFile); err != nil {
			return err
		}
		// set the permissions of the public key file
		if err := os.Chmod(publicKeyFile.Path, 0644); err != nil {
			return err
		}

	}

	return nil
}
