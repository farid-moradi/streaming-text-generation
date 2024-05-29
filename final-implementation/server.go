package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const text = `
This repository contains code examples and an implementation for building a streaming text generation interface using Server-Sent Events (SSE) in Golang. SSE is a lightweight protocol for real-time communication over HTTP, ideal for scenarios like Large Language Model (LLM) interfaces.
`

// ExpRand generates a random value from an exponential distribution
// with the given rate parameter (lambda) and location shift (mu).
func ExpRand(lambda, mu float64) float64 {
	// Generate a random value from the standard exponential distribution
	u := rand.ExpFloat64()

	// Apply the rate parameter (lambda)
	x := u / lambda

	// Add the location shift (mu)
	return x + mu
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Replace with your allowed origins
	}))

	// Define a route to handle the incremental text streaming
	e.GET("/stream", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		words := strings.Fields(text)
		for _, word := range words {
			fmt.Fprintf(c.Response(), "data: %s\n\n", word)
			c.Response().Flush()
			time.Sleep(time.Duration(ExpRand(0.07, 100)) * time.Millisecond)
		}

		return nil
	})

	// Start the Echo server
	e.Start(":8080")
}
