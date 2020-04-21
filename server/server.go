package server

import (
	"github.com/DiscoreMe/minego/protocol/packet"
	"net"
	"sync"
)

type Server struct {
	l           net.Listener
	handlers    map[int]MiddlewareFunc
	muxHandlers sync.RWMutex
	ErrHandler  func(error)
}

func NewServer(l net.Listener) *Server {
	s := &Server{
		l:           l,
		handlers:    make(map[int]MiddlewareFunc),
		muxHandlers: sync.RWMutex{},
	}
	s.setDefaultHandlerErr()
	return s
}

func (s *Server) setDefaultHandlerErr() {
	s.ErrHandler = func(error) {}
}

func (s *Server) Listen() error {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			return err
		}
		client := NewClient(conn)
		_, err = client.PackLength()
		if err != nil {
			client.Disconnect()
			s.ErrHandler(err)
			continue
		}
		id, err := client.PacketID()
		if err != nil {
			client.Disconnect()
			s.ErrHandler(err)
			continue
		}
		go func() {
			s.handlerRequest(packet.FindPacketByID(id), client)
			client.Disconnect()
		}()
	}
}

func (s *Server) handlerRequest(p packet.Packet, c *Client) {
	s.muxHandlers.RLock()
	hr, ok := s.handlers[p.ID()]
	s.muxHandlers.RUnlock()

	if ok {
		err := hr(c)
		if err != nil {
			s.ErrHandler(err)
		}
	}
	// todo handle unknown handlers
}

type MiddlewareFunc func(c *Client) error

func (s *Server) HandleFunc(p packet.Packet, fn MiddlewareFunc) {
	s.handlers[p.ID()] = fn
}
