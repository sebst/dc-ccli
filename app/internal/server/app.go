/*
Copyright © 2024 devcontainer.com
*/
package server

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"time"

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

	app.Get("/logs", func(c *fiber.Ctx) error {
		// Open the log file for streaming
		fileName := "/tmp/logfile"
		file, err := os.Open(fileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to open " + fileName + ": " + err.Error())
		}
		// Do not defer file.Close() since we want to keep it open for streaming

		// Set headers for a streaming response
		c.Set("Content-Type", "text/plain")

		// Set the BodyStreamWriter to continuously stream data
		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			reader := bufio.NewReader(file)

			// Infinite loop to keep reading and sending new lines
			for {
				// Read one line from the file
				line, _, err := reader.ReadLine()

				if err != nil {
					// If we encounter EOF, sleep a bit and continue to wait for new data
					if err.Error() == "EOF" {
						time.Sleep(500 * time.Millisecond)
						continue
					}

					// If another error occurs (e.g., file read error), send the error and break out
					w.WriteString("Error reading file: " + err.Error() + "\n")
					w.Flush()
					return
				}

				// Write the line to the stream
				if len(line) > 0 {
					w.WriteString(string(line) + "\n")
					w.Flush() // Immediately flush the data to the client
				}
			}
		})

		return nil
	})

	app.Get("/api/processes", func(c *fiber.Ctx) error {
		// Get running processes
		processList, err := getRunningProcesses()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(processList)
	})

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

func RunApp(port int) {
	app := GetApp()
	listenArg := fmt.Sprintf(":%d", port)
	app.Listen(listenArg)
}
