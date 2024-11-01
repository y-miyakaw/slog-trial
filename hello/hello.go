package hello

import (
	"fmt"
	"net/http"
	"slog-trial/logger"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	fmt.Println("Hello, World! Hello")
	logger.GlobalLogger.Info("Hello, World! Hello")
	return c.String(http.StatusOK, "Hello, World! Hello")
}
