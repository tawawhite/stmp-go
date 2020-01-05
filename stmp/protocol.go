// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-23 19:55:29
package stmp

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
	"strings"
)

func ReadNegotiate(input string) (v string, n int) {
	n = strings.IndexByte(input, ',')
	if n == -1 {
		n = len(input)
	}
	seg := strings.IndexByte(input[0:n], ';')
	if seg == -1 {
		seg = n
	}
	v = strings.TrimSpace(input[0:seg])
	return
}

func EscapeHeadKey(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(key, "%", "%25"), ":", "%3A")
}

func EscapeHeadValue(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "%", "%25"), "\n", "%0A")
}

func UnescapeHeadKey(key string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(key, "%3A", ":"), "%25", "%"))
}

func UnescapeHeadValue(value string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(value, "%0A", "\n"), "%25", "%"))
}

func UvarintSize(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}

func ReadUvarint(r io.Reader, b1 []byte) (uint64, error) {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		_, err := r.Read(b1)
		if err != nil {
			return x, err
		}
		b := b1[0]
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return x, errors.New("uint64 overflow")
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

func ReadUint16(r io.Reader, b2 []byte) (v uint16, err error) {
	_, err = r.Read(b2)
	if err != nil {
		return
	}
	v = binary.LittleEndian.Uint16(b2)
	return
}

var digits = [36]byte{}
var chunks = [256]byte{}
var hexOffsets = [16]byte{0, 4, 8, 12, 16, 1}

func init() {
	for i := 0; i < 256; i++ {
		chunks[i] = 255
	}
	var i byte
	for i = '0'; i <= '9'; i++ {
		digits[i-'0'] = i
		chunks[i] = i - '0'
	}
	for i = 'A'; i <= 'Z'; i++ {
		digits[i-'A'+10] = i
		chunks[i] = i - 'A' + 10
	}
	for i = 'a'; i < 'z'; i++ {
		chunks[i] = i - 'a' + 10
	}
	for i = 0; i < 16; i++ {
		hexOffsets[i] = i * 4
	}
}

func AppendHex(u uint64, buf []byte) int {
	i := len(buf)
	for u > 15 {
		i--
		buf[i] = digits[u&0xF]
		u >>= 4
	}
	// u < base
	i--
	buf[i] = digits[u]
	copy(buf, buf[i:])
	return len(buf) - i
}

func ParseHexStatus(buf []byte) (s Status, err error) {
	m := len(buf)
	if m > 2 || m == 0 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if chunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		s |= Status(uint16(chunks[c]) << hexOffsets[m-i-1])
	}
	return
}

func ParseHexUint16(buf []byte) (n uint16, err error) {
	m := len(buf)
	if m > 4 || m == 0 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if chunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		n |= uint16(chunks[c]) << hexOffsets[m-i-1]
	}
	return n, nil
}

func ParseHexUint64(buf []byte) (n uint64, err error) {
	m := len(buf)
	if m == 0 || m > 16 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if chunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		n |= uint64(chunks[c]) << hexOffsets[m-i-1]
	}
	return n, nil
}