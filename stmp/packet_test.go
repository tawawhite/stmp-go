package stmp_test

import (
	"bytes"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func writePacket(p *stmp.Packet) ([]byte, stmp.StatusError) {
	sb := stmp.NewBuilder()
	se := p.Write(sb)
	return sb.Get(), se
}

type packetCase struct {
	name   string
	packet *stmp.Packet
	write  []byte
	binary []byte
	text   string
}

var h stmp.Header

func newHeader(k string) stmp.Header {
	h = stmp.NewHeader()
	h.Set(k, "case header\n")
	return h
}

var p []byte

func newPayload(prefix string) []byte {
	p = []byte(prefix + ":\nP;")
	return p
}

var packetCases = []packetCase{
	{
		name:   "ping",
		packet: &stmp.Packet{Kind: stmp.MessageKindPing},
		write:  []byte{0x80},
		text:   "I",
	},
	{
		name:   "pong",
		packet: &stmp.Packet{Kind: stmp.MessageKindPong},
		write:  []byte{0b10010000},
		text:   "O",
	},
	{
		name:   "close",
		packet: &stmp.Packet{Kind: stmp.MessageKindClose, Status: 0x01, WithPayload: true, Payload: newPayload("Close")},
		write:  stmp.NewBuilder().Byte(0b11011000).Byte(0x01).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b11011000).Byte(0x01).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("C1").PString("P").Bytes(p).Build(),
	},
	{
		name:   "empty close",
		packet: &stmp.Packet{Kind: stmp.MessageKindClose, Status: 0xF1},
		write:  stmp.NewBuilder().Byte(0b11010000).Byte(0xF1).Get(),
		binary: stmp.NewBuilder().Byte(0b11010000).Byte(0xF1).Get(),
		text:   stmp.NewBuilder().String("CF1").Build(),
	},
	{
		name:   "response",
		packet: &stmp.Packet{Kind: stmp.MessageKindResponse, Mid: 0x1234, Status: 0x01, WithHeader: true, Header: newHeader("Response"), WithPayload: true, Payload: newPayload("Response")},
		write:  stmp.NewBuilder().Byte(0b11001100, 0x34, 0x12).Byte(0x01).SBytes(h.Marshal()).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b11001100, 0x34, 0x12).Byte(0x01).SBytes(h.Marshal()).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("S1234:1").PString("H").Bytes(h.Marshal()).PString("P").Bytes(p).Build(),
	},
	{
		name:   "empty response",
		packet: &stmp.Packet{Kind: stmp.MessageKindResponse, Mid: 0x1234, Status: 0x01},
		write:  stmp.NewBuilder().Byte(0b11000000, 0x34, 0x12).Byte(0x01).Get(),
		binary: stmp.NewBuilder().Byte(0b11000000, 0x34, 0x12).Byte(0x01).Get(),
		text:   stmp.NewBuilder().String("S1234:1").Build(),
	},
	{
		name:   "without header response",
		packet: &stmp.Packet{Kind: stmp.MessageKindResponse, Mid: 0x1234, Status: 0x01, WithPayload: true, Payload: newPayload("Response")},
		write:  stmp.NewBuilder().Byte(0b11001000, 0x34, 0x12).Byte(0x01).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b11001000, 0x34, 0x12).Byte(0x01).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("S1234:1").PString("P").Bytes(p).Build(),
	},
	{
		name:   "without payload response",
		packet: &stmp.Packet{Kind: stmp.MessageKindResponse, Mid: 0x1234, Status: 0x01, WithHeader: true, Header: newHeader("Response")},
		write:  stmp.NewBuilder().Byte(0b11000100, 0x34, 0x12).Byte(0x01).SBytes(h.Marshal()).Get(),
		binary: stmp.NewBuilder().Byte(0b11000100, 0x34, 0x12).Byte(0x01).SBytes(h.Marshal()).Get(),
		text:   stmp.NewBuilder().String("S1234:1").PString("H").Bytes(h.Marshal()).Build(),
	},
	{
		name:   "action request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: false, WithPayload: true, WithHeader: true, Action: 0xFFFF, Header: newHeader("Action-Request"), Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10101100, 0x34, 0x12).Uvarint(0xFFFF).SBytes(h.Marshal()).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10101100, 0x34, 0x12).Uvarint(0xFFFF).SBytes(h.Marshal()).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("Q1234:FFFF").PString("H").Bytes(h.Marshal()).PString("P").Bytes(p).Build(),
	},
	{
		name:   "method request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: true, WithPayload: true, WithHeader: true, Method: "stmp.test.packet.All", Header: newHeader("Action-Request"), Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10101110, 0x34, 0x12).SString("stmp.test.packet.All").SBytes(h.Marshal()).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10101110, 0x34, 0x12).SString("stmp.test.packet.All").SBytes(h.Marshal()).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("Q1234:M:stmp.test.packet.All").PString("H").Bytes(h.Marshal()).PString("P").Bytes(p).Build(),
	},
	{
		name:   "without payload action request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: false, WithHeader: true, Action: 0xFFFF, Header: newHeader("Action-Request")},
		write:  stmp.NewBuilder().Byte(0b10100100, 0x34, 0x12).Uvarint(0xFFFF).SBytes(h.Marshal()).Get(),
		binary: stmp.NewBuilder().Byte(0b10100100, 0x34, 0x12).Uvarint(0xFFFF).SBytes(h.Marshal()).Get(),
		text:   stmp.NewBuilder().String("Q1234:FFFF").PString("H").Bytes(h.Marshal()).Build(),
	},
	{
		name:   "without payload method request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: true, WithHeader: true, Method: "stmp.test.packet.All", Header: newHeader("Action-Request")},
		write:  stmp.NewBuilder().Byte(0b10100110, 0x34, 0x12).SString("stmp.test.packet.All").SBytes(h.Marshal()).Get(),
		binary: stmp.NewBuilder().Byte(0b10100110, 0x34, 0x12).SString("stmp.test.packet.All").SBytes(h.Marshal()).Get(),
		text:   stmp.NewBuilder().String("Q1234:M:stmp.test.packet.All").PString("H").Bytes(h.Marshal()).Build(),
	},
	{
		name:   "without header action request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: false, WithPayload: true, Action: 0xFFFF, Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10101000, 0x34, 0x12).Uvarint(0xFFFF).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10101000, 0x34, 0x12).Uvarint(0xFFFF).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("Q1234:FFFF").PString("P").Bytes(p).Build(),
	}, {
		name:   "without header method request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: true, WithPayload: true, Method: "stmp.test.packet.All", Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10101010, 0x34, 0x12).SString("stmp.test.packet.All").SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10101010, 0x34, 0x12).SString("stmp.test.packet.All").Bytes(p).Get(),
		text:   stmp.NewBuilder().String("Q1234:M:stmp.test.packet.All").PString("P").Bytes(p).Build(),
	},
	{
		name:   "without header without payload action request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: false, Action: 0xFFFF},
		write:  stmp.NewBuilder().Byte(0b10100000, 0x34, 0x12).Uvarint(0xFFFF).Get(),
		binary: stmp.NewBuilder().Byte(0b10100000, 0x34, 0x12).Uvarint(0xFFFF).Get(),
		text:   stmp.NewBuilder().String("Q1234:FFFF").Build(),
	},
	{
		name:   "without header without payload method request",
		packet: &stmp.Packet{Kind: stmp.MessageKindRequest, Mid: 0x1234, StringAction: true, Method: "stmp.test.packet.All"},
		write:  stmp.NewBuilder().Byte(0b10100010, 0x34, 0x12).SString("stmp.test.packet.All").Get(),
		binary: stmp.NewBuilder().Byte(0b10100010, 0x34, 0x12).SString("stmp.test.packet.All").Get(),
		text:   stmp.NewBuilder().String("Q1234:M:stmp.test.packet.All").Build(),
	},
	{
		name:   "action notify",
		packet: &stmp.Packet{Kind: stmp.MessageKindNotify, Mid: 0, StringAction: false, WithPayload: true, WithHeader: true, Action: 0xFFFF, Header: newHeader("Action-Request"), Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10111100).Uvarint(0xFFFF).SBytes(h.Marshal()).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10111100).Uvarint(0xFFFF).SBytes(h.Marshal()).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("NFFFF").PString("H").Bytes(h.Marshal()).PString("P").Bytes(p).Build(),
	},
	{
		name:   "method notify",
		packet: &stmp.Packet{Kind: stmp.MessageKindNotify, Mid: 0, StringAction: true, WithPayload: true, WithHeader: true, Method: "stmp.test.packet.All", Header: newHeader("Action-Request"), Payload: newPayload("action request")},
		write:  stmp.NewBuilder().Byte(0b10111110).SString("stmp.test.packet.All").SBytes(h.Marshal()).SBytes(p).Get(),
		binary: stmp.NewBuilder().Byte(0b10111110).SString("stmp.test.packet.All").SBytes(h.Marshal()).Bytes(p).Get(),
		text:   stmp.NewBuilder().String("NM:stmp.test.packet.All").PString("H").Bytes(h.Marshal()).PString("P").Bytes(p).Build(),
	},
}

func TestPacket(t *testing.T) {
	for _, c := range packetCases {
		if c.binary == nil {
			c.binary = c.write
		}
		w, se := writePacket(c.packet)
		assert.Nil(t, se, c.name+": write error")
		assert.Equal(t, c.write, w, c.name+": write")
		assert.Equal(t, c.binary, c.packet.MarshalBinary(), c.name+": marshal binary")
		assert.Equal(t, []byte(c.text), c.packet.MarshalText(), c.name+": marshal text")
		p := new(stmp.Packet)
		se = p.Read(bytes.NewReader(c.write), 0xFFFFFF)
		assert.Nil(t, se, c.name+": read error")
		assert.Equal(t, c.packet, p, c.name+": read")
		p = new(stmp.Packet)
		se = p.UnmarshalBinary(c.binary)
		assert.Nil(t, se, c.name+": unmarshal binary error")
		assert.Equal(t, c.packet, p, c.name+": unmarshal binary")
		p = new(stmp.Packet)
		se = p.UnmarshalText([]byte(c.text))
		assert.Nil(t, se, c.name+": unmarshal text error")
		assert.Equal(t, c.packet, p, c.name+": unmarshal text")
	}
}
