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
	v = strings.TrimSpace(v[0:seg])
	return
}

func escapeHeadKey(key string) string {
	return strings.ReplaceAll(strings.ReplaceAll(key, "%", "%25"), ":", "%3A")
}

func escapeHeadValue(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "%", "%25"), "\n", "%0A")
}

func unescapeHeadKey(key string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(key, "%3A", ":"), "%25", "%"))
}

func unescapeHeadValue(value string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(value, "%0A", "\n"), "%25", "%"))
}

func varintLen(x uint64) (n int) {
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
