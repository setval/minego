package core

import (
	"encoding/json"
	"fmt"
	"github.com/DiscoreMe/minego/protocol/codec"
	"github.com/DiscoreMe/minego/protocol/packet"
	"github.com/DiscoreMe/minego/server"
)

func HandleHandshake(c *server.Client) error {
	var p packet.Handshake
	if err := c.Decode(&p); err != nil {
		return err
	}

	fmt.Printf("Protocol Version: %d\nServer Address: %s\nServer Port: %d\nNext state: %d\n\n", p.ProtoVersion, p.ServerAddress, p.ServerPort, p.NextState)

	if p.NextState != 1 {
		return nil
	}

	type Players struct {
		Max    int        `json:"max"`
		Online int        `json:"online"`
		Sample []struct{} `json:"sample"`
	}

	type Data struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Description struct {
			Text string `json:"text"`
		} `json:"description"`
		Players Players `json:"players"`
		Favicon string  `json:"favicon"`
	}

	data := Data{}
	data.Version.Protocol = 578
	data.Version.Name = "Test Server"
	data.Description.Text = "Desc"
	data.Players = Players{
		Max:    10,
		Online: 0,
		Sample: []struct{}{},
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	m := codec.String(string(b))

	fmt.Println(string(b))

	if err := c.Write(&packet.Handshake{}, &m); err != nil {
		return err
	}

	return nil
}

func LegacyHandleHandshake(c *server.Client) error {
	var p packet.LegacyHandshaking
	if err := c.Decode(&p); err != nil {
		return err
	}

	return nil
}
