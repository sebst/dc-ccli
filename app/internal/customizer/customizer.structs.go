/*
Copyright © 2024 devcontainer.com
*/
package customizer

type Config struct {
	Dotfiles    Dotfiles            `json:"dotfiles"`
	Environment map[string]string   `json:"environment"`
	Files       map[string]File     `json:"files"`
	Packages    []Package           `json:"packages"`
	Vscode      map[string]struct{} `json:"vscode"`
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
	Content     Content `json:"content"`
	Permissions string  `json:"permissions"`
	Owner       Owner   `json:"owner"`
	Group       Group   `json:"group"`
}

type Content struct {
	Plain string `json:"plain"`
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
