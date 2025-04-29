package main

import (
	"github.com/koki-noguchi/websocket-practice/app/service"
	"github.com/koki-noguchi/websocket-practice/handler"
	"github.com/koki-noguchi/websocket-practice/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
)

func main() {
	logger.Init()
	defer logger.S().Sync()

	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		logger.S().Info("hello world")
		return c.String(http.StatusOK, "Hello, World!")
	})
	roomService := service.NewRoomService()
	webSocketHandler := handler.NewWebSocketHandler(roomService)

	e.GET("/ws", webSocketHandler.HandleWebsocket)

	if err := e.Start(":8080"); err != nil {
		slog.Error(err.Error())
	}
}
