package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/koki-noguchi/websocket-practice/app/model"
	"github.com/koki-noguchi/websocket-practice/app/service"
	"github.com/koki-noguchi/websocket-practice/helper"
	"github.com/koki-noguchi/websocket-practice/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebSocketHandler struct {
	RoomService service.RoomServiceInterface
}

type Message struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketHandler(roomService service.RoomServiceInterface) *WebSocketHandler {
	return &WebSocketHandler{
		RoomService: roomService,
	}
}

func (h *WebSocketHandler) HandleWebsocket(c echo.Context) error {
	// websocketに昇格
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "WebSocket Upgrade failed: "+err.Error())
	}
	defer ws.Close()

	// クライアントを定義
	clientId := uuid.New().String()
	client := model.NewClient(ws, clientId)

	_, roomNameByte, err := ws.ReadMessage()
	if err != nil {
		logger.S().Info("connection closed during room name read: " + err.Error())
		return err
	}
	roomName := string(roomNameByte)

	room := h.RoomService.GetOrCreateRoom(roomName)
	room.Join(client)

	defer room.Leave(client)
	// クライアントにメッセージを送信
	go writePump(client)

	// クライアントからのメッセージ受信
	// ブロードキャストに入れる
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				logger.S().Info("context canceled")
			} else if errors.Is(err, context.DeadlineExceeded) {
				logger.S().Warn("context deadline exceeded")
			} else if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logger.S().Info("connection closed")
			} else {
				logger.S().Error("read error: " + err.Error())
			}
			break
		}
		room.Broadcast <- model.BroadcastMessage{
			Sender:  client,
			Message: helper.MustMarshal(Message{Text: string(message)}),
		}
	}
	return nil
}

func writePump(c *model.Client) {
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			logger.S().Error("write error: " + err.Error())
			break
		}
	}
}
