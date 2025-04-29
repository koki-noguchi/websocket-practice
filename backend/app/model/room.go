package model

import "sync"

type Room struct {
	Name      string
	Broadcast chan BroadcastMessage
	Clients   map[*Client]bool
	mu        sync.RWMutex
}

type BroadcastMessage struct {
	Sender  *Client
	Message []byte
}

func NewRoom(roomName string) *Room {
	return &Room{
		Name:      roomName,
		Broadcast: make(chan BroadcastMessage),
		Clients:   make(map[*Client]bool),
	}
}

func (r *Room) Start() {
	for {
		b := <-r.Broadcast
		r.mu.Lock()
		for client := range r.Clients {
			select {
			case client.Send <- b.Message:
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
