package stmp_test

import (
	"github.com/acrazing/stmp-go/stmp"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type handshakeCase struct {
	name   string
	client bool
	packet *stmp.Handshake
	binary []byte
	text   []byte
}

var hh stmp.Header

func newHandshakeHeader(k string) stmp.Header {
	hh = stmp.NewHeader()
	hh.Set(k, "case header\n")
	return hh
}

var m string

func newMessage(name string) string {
	m = name
	return m
}

var handshakeCases = []handshakeCase{
	{
		name:   "client handshake",
		packet: &stmp.Handshake{Kind: stmp.HandshakeKindClient, Major: 0xF, Minor: 0xF, Header: newHandshakeHeader("client handshake header"), Message: newMessage("client handshake message")},
		binary: stmp.NewBuilder().String("STMP").Byte(0xFF).SBytes(hh.Marshal()).SString(m).Get(),
		text:   stmp.NewBuilder().String("STMP").String("FF").PString("H").Bytes(hh.Marshal()).PString("M").String(m).Get(),
	},
	{
		name:   "server handshake",
		packet: &stmp.Handshake{Kind: stmp.HandshakeKindServer, Status: 0xF1, Header: newHandshakeHeader("server handshake header"), Message: newMessage("server handshake message")},
		binary: stmp.NewBuilder().String("STMP").Byte(0xF1).SBytes(hh.Marshal()).SString(m).Get(),
		text:   stmp.NewBuilder().String("STMP").String("F1").PString("H").Bytes(hh.Marshal()).PString("M").String(m).Get(),
	},
}

func TestHandshake(t *testing.T) {
	for _, c := range handshakeCases {
		log.Printf("packet header: %q.", c.packet.Header.Marshal())
		assert.Equal(t, c.binary, c.packet.MarshalBinary(), c.name+": marshal binary")
		assert.Equal(t, c.text, c.packet.MarshalText(), c.name+": marshal text")
		p := &stmp.Handshake{Kind: c.packet.Kind, Header: stmp.NewHeader()}
		se := p.UnmarshalBinary(c.binary)
		assert.Nil(t, se, c.name+": unmarshal binary error")
		assert.Equal(t, c.packet, p, c.name+": unmarshal binary")
		p = &stmp.Handshake{Kind: c.packet.Kind, Header: stmp.NewHeader()}
		se = p.UnmarshalText(c.text)
		assert.Nil(t, se, c.name+": unmarshal text error")
		assert.Equal(t, c.packet, p, c.name+": unmarshal text")
	}
}
