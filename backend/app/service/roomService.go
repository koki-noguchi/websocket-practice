package service

import (
	"github.com/koki-noguchi/websocket-practice/app/model"
	"sync"
)

type RoomService struct {
	rooms map[string]*model.Room
	mu    sync.Mutex
}

func NewRoomService() *RoomService {
	return &RoomService{
		rooms: make(map[string]*model.Room),
	}
}

func (s *RoomService) GetOrCreateRoom(name string) *model.Room {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, exists := s.rooms[name]
	if !exists {
		room = model.NewRoom(name)
		s.rooms[name] = room
		go room.Start()
	}

	return room
}
