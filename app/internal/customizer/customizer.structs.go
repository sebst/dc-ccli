/*
Copyright © 2024 devcontainer.com
*/
package customizer

type Config struct {
	Dotfiles    Dotfiles      `json:"dotfiles"`
	SshKeys     []SshKey      `json:"sshKeys"`
	Environment []Environment `json:"environment"`
	Files       []File        `json:"files"`
	// Files       map[string]File     `json:"files"`
	// Packages    []Package           `json:"packages"`
	// Vscode      map[string]struct{} `json:"vscode"`
}

type Environment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SshKey struct {
	PrivateKey string `json:"private"`
	PublicKey  string `json:"public"`
}

type Dotfiles struct {
	Github Github `json:"github"`
}

type Github struct {
	Repo    string `json:"repo"`
	Branch  string `json:"branch"`
	Path    string `json:"path"`
	Install string `json:"install"`
}

type File struct {
	// Content     Content `json:"content"`
	// Permissions string  `json:"permissions"`
	// Owner       Owner   `json:"owner"`
	// Group       Group   `json:"group"`
	Path    string  `json:"path"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

type Owner struct {
	UID int `json:"uid"`
}

type Group struct {
	GID int `json:"gid"`
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Manager string `json:"manager"`
}
