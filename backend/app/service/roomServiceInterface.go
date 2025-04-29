package service

import "github.com/koki-noguchi/websocket-practice/app/model"

type RoomServiceInterface interface {
	GetOrCreateRoom(name string) *model.Room
}
