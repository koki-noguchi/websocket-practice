package model

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Room *Room
	Id   string
}

func NewClient(conn *websocket.Conn, id string) *Client {
	return &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
		Id:   id,
	}
}
