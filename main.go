package main

import (
	"github.com/koki-noguchi/websocket-practice/handler"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", handler.HandleWebsocket)
	go handler.HandleMessage()

	if err := e.Start(":8080"); err != nil {
		slog.Error(err.Error())
	}
}
