/*
Copyright © 2024 devcontainer.com
*/
package server

import (
	"bufio"
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetApp() *fiber.App {
	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/app")
	})

	app.Get("/app/*", func(c *fiber.Ctx) error {
		// Return `indexHTML` as string with `text/html` content type
		return c.Type("html").Send(GetIndexHtml())
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Get("/api/processes", func(c *fiber.Ctx) error {
	// 	// Get running processes
	// 	processList, err := getRunningProcesses()
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": err.Error(),
	// 		})
	// 	}
	// 	return c.JSON(processList)
	// })

	// New route for streaming stdout of /bin/test.sh
	app.Get("/api/test", func(c *fiber.Ctx) error {
		cmd := exec.Command("/workspaces/ab-devcontainer-golang/api/test.sh")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		if err := cmd.Start(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		scanner := bufio.NewScanner(stdout)
		c.Set("Content-Type", "text/plain")
		c.Set(fiber.HeaderContentEncoding, "chunked")
		c.Set(fiber.HeaderTransferEncoding, "chunked")
		c.Set(fiber.HeaderConnection, "keep-alive")

		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			for scanner.Scan() {
				text := scanner.Text()
				if _, err := w.WriteString(text + "\n"); err != nil {
					fmt.Println("Error writing:", err)
					return
				}
				fmt.Println(text)
				if err := w.Flush(); err != nil {
					fmt.Println("Error flushing:", err)
					return
				}
				w.Flush()
			}
			w.Flush()
			cmd.Wait()
		})

		return nil
	})

	return app

	// Start the server on port 3000
	// app.Listen(":3001")
}
