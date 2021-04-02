package main

import (
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v2"
	"strings"
	"sync"
)

type User struct {
	ID string
	UUID string
	PhotoURL string

	Connection *websocket.Conn

	OutBuffer chan []byte
	PeerConnection *webrtc.PeerConnection

	InTrack sync.Map
	OutTrack sync.Map

	IsMuted bool
	CloseChan chan bool
}

func NewUser(uuid string) *User {
	return &User{
		ID:             "",
		UUID:           uuid,
		PhotoURL:       photoUrlFromUuid(uuid),
		Connection:     nil,
		OutBuffer:      nil,
		PeerConnection: nil,
		InTrack:        sync.Map{},
		OutTrack:       sync.Map{},
		IsMuted:        false,
		CloseChan:      nil,
	}
}

func photoUrlFromUuid(uuid string) string {
	trimmedUuid := strings.ReplaceAll(uuid, "-", "")
	return "https://crafatar.com/avatars/" + trimmedUuid
}
