package main

import "time"

type Server struct {
	Name string
	Users map[string]*User

	Volumes map[*User]uint8
	VolumeChan chan *volume

	JoinRequests chan *User
	LeaveRequests chan *User

	CloseChan chan bool

	AliveTimer *time.Ticker
}

type volume struct {

}

func NewServer(name string) *Server {
	return &Server{
		Name:          name,
		Users:         make(map[string]*User),
		Volumes:       make(map[*User]uint8),
		JoinRequests:  make(chan *User),
		LeaveRequests: make(chan *User),
		CloseChan: make(chan bool),
		AliveTimer: time.NewTicker(1 * time.Minute),
	}
}

func (s *Server) Join(user *User) {
	s.JoinRequests <- user
}

func (s *Server) Leave(user *User) {
	s.LeaveRequests <- user
}

func (s *Server) Run() {
	go func() {
		for {
			select {
			case user := <- s.JoinRequests:
				s.doJoin(user)
			case user := <- s.LeaveRequests:
				s.doLeave(user)
			case <- s.AliveTimer.C:
				s.doAliveTest()
			case <- s.CloseChan:
				s.close()
				break
			}
		}
	}()
}

func (s *Server) doJoin(user *User) {

}

func (s *Server) doLeave(user *User) {

}

func (s *Server) close() {
	close(s.JoinRequests)
	close(s.LeaveRequests)
	s.AliveTimer.Stop()
}

func (s *Server) doAliveTest() {

}
