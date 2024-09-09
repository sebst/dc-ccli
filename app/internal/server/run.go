/*
Copyright © 2024 devcontainer.com
*/
package server

import "github.com/gofiber/fiber/v2"

func RunServer(app *fiber.App) {
	app.Listen(":8080")
}
