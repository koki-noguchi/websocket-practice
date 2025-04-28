package model

type Room struct {
	name      string
	broadcast chan []byte
	clients   map[*Client]bool
}
