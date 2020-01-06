package stmp

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
	"strings"
)

// read input value
func readNegotiate(input string) (v string, n int) {
	n = strings.IndexByte(input, ',')
	if n == -1 {
		n = len(input)
	} else {
		n += 1
	}
	seg := strings.IndexByte(input[0:n], ';')
	if seg == -1 {
		seg = n
	}
	v = strings.TrimSpace(input[0:seg])
	return
}

func uvarintSize(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}

func readUvarint(r io.Reader, b1 []byte) (uint64, error) {
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

func readUint16(r io.Reader, b2 []byte) (v uint16, err error) {
	_, err = r.Read(b2)
	if err != nil {
		return
	}
	v = binary.LittleEndian.Uint16(b2)
	return
}

var hexDigits = [36]byte{}
var hexChunks = [256]byte{}
var hexOffsets = [16]byte{0, 4, 8, 12, 16, 1}
var hexCaches [256]string

func init() {
	for i := 0; i < 256; i++ {
		hexChunks[i] = 255
	}
	var i byte
	for i = '0'; i <= '9'; i++ {
		hexDigits[i-'0'] = i
		hexChunks[i] = i - '0'
	}
	for i = 'A'; i <= 'Z'; i++ {
		hexDigits[i-'A'+10] = i
		hexChunks[i] = i - 'A' + 10
	}
	for i = 'a'; i < 'z'; i++ {
		hexChunks[i] = i - 'a' + 10
	}
	for i = 0; i < 16; i++ {
		hexOffsets[i] = i * 4
		hexCaches[i] = string(hexDigits[i])
	}
	for i := 0x10; i < 0x100; i++ {
		hexCaches[i] = string(hexDigits[i>>4]) + string(hexDigits[i&0xF])
	}
}

func hexAppend(u uint64, buf []byte) int {
	i := len(buf)
	for u > 15 {
		i--
		buf[i] = hexDigits[u&0xF]
		u >>= 4
	}
	// u < base
	i--
	buf[i] = hexDigits[u]
	copy(buf, buf[i:])
	return len(buf) - i
}

func hexFormatUint64(u uint64) string {
	buf := make([]byte, 8, 8)
	i := 8
	for u > 15 {
		i--
		buf[i] = hexDigits[u&0xF]
		u >>= 4
	}
	i--
	buf[i] = hexDigits[u]
	return string(buf[i:])
}

func hexParseStatus(buf []byte) (s Status, err error) {
	m := len(buf)
	if m > 2 || m == 0 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if hexChunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		s |= Status(uint16(hexChunks[c]) << hexOffsets[m-i-1])
	}
	return
}

func hexParseUint16(buf []byte) (n uint16, err error) {
	m := len(buf)
	if m > 4 || m == 0 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if hexChunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		n |= uint16(hexChunks[c]) << hexOffsets[m-i-1]
	}
	return n, nil
}

func hexParseUint64(buf []byte) (n uint64, err error) {
	m := len(buf)
	if m == 0 || m > 16 {
		err = errors.New("out of range")
		return
	}
	for i, c := range buf {
		if hexChunks[c] > 15 {
			err = errors.New("invalid bit: " + string(c))
			return
		}
		n |= uint64(hexChunks[c]) << hexOffsets[m-i-1]
	}
	return n, nil
}
