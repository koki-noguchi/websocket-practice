package model

import "sync"

type Room struct {
	name      string
	broadcast chan []byte
	clients   map[*Client]bool
	mu        sync.RWMutex
}

func NewRoom(roomName string) *Room {
	return &Room{
		name:      roomName,
		broadcast: make(chan []byte),
		clients:   make(map[*Client]bool),
	}
}

func (r *Room) Start() {
	for {
		msg := <-r.broadcast
		r.mu.Lock()
		for client := range r.clients {
			select {
			case client.send <- msg:
			default:
				delete(r.clients, client)
				close(client.send)
			}
		}
		r.mu.Unlock()
	}
}
