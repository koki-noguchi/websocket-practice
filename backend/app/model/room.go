package model

import "sync"

type Room struct {
	Name      string
	Broadcast chan []byte
	Clients   map[*Client]bool
	mu        sync.RWMutex
}

func NewRoom(roomName string) *Room {
	return &Room{
		Name:      roomName,
		Broadcast: make(chan []byte),
		Clients:   make(map[*Client]bool),
	}
}

func (r *Room) Start() {
	for {
		msg := <-r.Broadcast
		r.mu.Lock()
		for client := range r.Clients {
			select {
			case client.Send <- msg:
			default:
				delete(r.Clients, client)
				close(client.Send)
			}
		}
		r.mu.Unlock()
	}
}

func (r *Room) Join(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Clients[client] = true
	client.Room = r
}

func (r *Room) Leave(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Clients, client)
	close(client.Send)
}
