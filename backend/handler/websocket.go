package handler

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/koki-noguchi/websocket-practice/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{}
var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)
var mu sync.Mutex

func HandleWebsocket(c echo.Context) error {
	// origin check
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// websocketに昇格
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// クライアントを定義
	client := &Client{conn: ws, send: make(chan []byte, 256)}
	mu.Lock()
	clients[client] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, client)
		mu.Unlock()
		close(client.send)
	}()

	// クライアントにメッセージを送信
	go func(c *Client) {
		for msg := range c.send {
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.S().Error("write error: " + err.Error())
				break
			}
		}
	}(client)

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
		broadcast <- message
	}
	return nil
}

func HandleMessage() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
