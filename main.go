package main

import (
	"fmt"
	"github.com/DiscoreMe/minego/core"
	"github.com/DiscoreMe/minego/protocol/packet"
	"github.com/DiscoreMe/minego/server"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		panic(err)
	}

	serv := server.NewServer(ln)

	serv.ErrHandler = func(err error) {
		fmt.Println("err: ", err)
	}

	serv.HandleFunc(&packet.Handshake{}, core.HandlerHandshake)

	if err := serv.Listen(); err != nil {
		panic(err)
	}
}

func handshakeFunc() error {
	fmt.Println("call handshakeFunc")
	return nil
}
