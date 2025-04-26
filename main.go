package main

import (
	"github.com/koki-noguchi/websocket-practice/handler"
	"github.com/koki-noguchi/websocket-practice/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
)

func main() {
	logger.Init()
	defer logger.SugaredLogger.Sync()

	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		logger.S().Info("hello world")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", handler.HandleWebsocket)
	go handler.HandleMessage()

	if err := e.Start(":8080"); err != nil {
		slog.Error(err.Error())
	}
}
