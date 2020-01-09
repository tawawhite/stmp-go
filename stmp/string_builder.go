// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-09 15:36:00
package stmp

import (
	"encoding/binary"
)

type Builder struct {
	buf []byte
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (sb *Builder) Uvarint(v uint64) *Builder {
	chunk := make([]byte, 10)
	n := binary.PutUvarint(chunk, v)
	sb.buf = append(sb.buf, chunk[:n]...)
	return sb
}

func (sb *Builder) Byte(b ...byte) *Builder {
	sb.buf = append(sb.buf, b...)
	return sb
}

func (sb *Builder) SBytes(bs ...[]byte) *Builder {
	for _, c := range bs {
		sb.Uvarint(uint64(len(c))).Bytes(c)
	}
	return sb
}

func (sb *Builder) SString(s ...string) *Builder {
	for _, c := range s {
		sb.Uvarint(uint64(len(c))).String(c)
	}
	return sb
}

func (sb *Builder) PBytes(bs ...[]byte) *Builder {
	for _, c := range bs {
		sb.Byte('\n').Bytes(c)
	}
	return sb
}

func (sb *Builder) PString(s ...string) *Builder {
	for _, c := range s {
		sb.Byte('\n').String(c)
	}
	return sb
}

func (sb *Builder) String(s ...string) *Builder {
	for _, c := range s {
		sb.buf = append(sb.buf, c...)
	}
	return sb
}

func (sb *Builder) Bytes(bs ...[]byte) *Builder {
	for _, c := range bs {
		sb.buf = append(sb.buf, c...)
	}
	return sb
}

func (sb *Builder) Write(data []byte) (int, error) {
	sb.buf = append(sb.buf, data...)
	return len(data), nil
}

func (sb *Builder) Flush() error {
	return nil
}

func (sb *Builder) Close() error {
	return nil
}

func (sb *Builder) Build() string {
	return string(sb.buf)
}

func (sb *Builder) Get() []byte {
	return sb.buf
}
