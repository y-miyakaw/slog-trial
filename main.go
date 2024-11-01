package main

import (
	"net/http"
	"slog-trial/hello"
	"slog-trial/logger"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		logger.GlobalLogger.Info("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", hello.Hello)
	e.Start(":8078")
}
