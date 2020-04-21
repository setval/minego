package main

import (
	"fmt"
	"github.com/DiscoreMe/minego/protocol/codec"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	var servPort codec.UShort
	if err := servPort.Decode(conn); err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	var protoVersion codec.VarInt
	if err := protoVersion.Decode(conn); err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	var servAddress codec.String
	if err := servAddress.Decode(conn); err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Printf("Protocol Version: %d\nServer Address: %s\nServer Port: %d\n", protoVersion, servAddress, servPort)
}
