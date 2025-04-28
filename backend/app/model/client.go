package model

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
	send chan []byte
	room *Room
	id   string
}
