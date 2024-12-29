/*
Copyright © 2024 devcontainer.com
*/
package socket

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type Message struct {
	UUID    string            `json:"uuid"`
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

type Result struct {
	UUID       string            `json:"uuid"`
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func ConnectSocket() {
	ctx := context.Background()
	endpoint := "ws://abc.devcontainer.online/ws"
	localEndpoint := "http://localhost:8000"

	for {
		err := connectAndHandle(ctx, endpoint, localEndpoint)
		if err != nil {
			fmt.Println("Connection closed, reconnecting...", err)
			time.Sleep(2 * time.Second)
		}
	}
}

func connectAndHandle(ctx context.Context, endpoint string, localEndpoint string) error {
	conn, _, err := websocket.Dial(ctx, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close(websocket.StatusInternalError, "the connection closed abnormally")

	fmt.Println("Connected to server")

	for {
		var message Message
		err := wsjson.Read(ctx, conn, &message)
		if err != nil {
			return fmt.Errorf("error reading message: %w", err)
		}

		go func(msg Message) {
			result, err := handleRequest(msg, localEndpoint)
			if err != nil {
				fmt.Println("Error handling request:", err)
				return
			}

			err = wsjson.Write(ctx, conn, result)
			if err != nil {
				fmt.Println("Error sending result:", err)
			}
		}(message)
	}
}

func handleRequest(message Message, localEndpoint string) (Result, error) {
	url := localEndpoint + message.URL

	// Prepare the HTTP request
	req, err := http.NewRequest(message.Method, url, nil)
	if err != nil {
		return Result{}, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range message.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Result{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, fmt.Errorf("failed to read response body: %w", err)
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = values[0]
	}

	fmt.Println("Processed Request for", message.Method, message.URL)

	return Result{
		UUID:       message.UUID,
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Body:       string(body),
	}, nil
}
