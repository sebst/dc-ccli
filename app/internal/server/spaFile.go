/*
Copyright © 2024 devcontainer.com
*/
package server

import (
	_ "embed"
)

//go:embed index.html
var indexHTML []byte

func GetIndexHtml() []byte {
	return indexHTML
}
