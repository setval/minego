package server

import (
	"fmt"
	"github.com/DiscoreMe/minego/protocol/codec"
	"github.com/DiscoreMe/minego/protocol/packet"
	"net"
	"sync"
)

type Server struct {
	l           net.Listener
	handlers    map[codec.VarInt]MiddlewareFunc
	muxHandlers sync.RWMutex
	ErrHandler  func(error)
}

func NewServer(l net.Listener) *Server {
	s := &Server{
		l:           l,
		handlers:    make(map[codec.VarInt]MiddlewareFunc),
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
		length, err := client.PackLength()
		if err != nil {
			client.Disconnect()
			s.ErrHandler(err)
			continue
		}
		if int32(length) < 0 {
			// todo constant
			s.ErrHandler(fmt.Errorf("packet length was too small, got %d", length))
			continue
		}
		if int32(length) < 1 {
			fmt.Println("legth < 1")
			continue
		}
		id, err := client.PacketID()
		if err != nil {
			client.Disconnect()
			s.ErrHandler(err)
			continue
		}

		if id < 0x00 {
			s.ErrHandler(fmt.Errorf("packet type was too small, got %d", id))
			continue
		}

		fmt.Printf("[connect] package %d has %d bytes\n", id, length)
		go func() {
			s.handlerRequest(packet.FindPacketByID(id), client)
		}()
	}
}

func (s *Server) handlerRequest(p packet.Packet, c *Client) {
	//defer c.Disconnect()

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
