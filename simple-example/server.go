package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Replace with your allowed origins
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Define a route to handle the incremental text streaming
	e.GET("/stream", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		var wordList = []string{
			"LLM", "Interface",
		}

		// Send text incrementally every second
		for i := 0; i < len(wordList); i++ {
			fmt.Fprintf(c.Response(), "data: %s\n\n", wordList[i])
			c.Response().Flush()
			time.Sleep(time.Second)
		}

		return nil
	})

	// Start the Echo server
	e.Start(":8080")
}
