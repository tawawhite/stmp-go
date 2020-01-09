// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-09 22:53:05

// base 64 number format
package num

import "errors"

const Digits = "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ._"

var bytes = [256]byte{}

func init() {
	for i := 0; i < 256; i++ {
		bytes[i] = 255
	}
	for i, c := range Digits {
		bytes[c] = byte(i)
	}
}

var (
	ErrEmpty    = errors.New("empty string")
	ErrInvalid  = errors.New("invalid string")
	ErrOverflow = errors.New("overflow")
)

func parse(s string) (uint64, error) {
	l := len(s)
	var n uint64
	var i = 0
	if l == 11 {
		n = uint64(bytes[s[0]])
		i = 1
	}
	for ; i < l; i++ {
		if bytes[s[i]] == 0xFF {
			return 0, ErrInvalid
		}
		n = n<<6 + uint64(bytes[s[i]])
	}
	return n, nil
}

func ParseUint32(s string) (uint32, error) {
	l := len(s)
	if l == 0 {
		return 0, ErrEmpty
	}
	if l > 6 || (l == 6 && bytes[s[0]] >= 4) {
		return 0, ErrOverflow
	}
	n, err := parse(s)
	return uint32(n), err
}

func ParseUint64(s string) (uint64, error) {
	l := len(s)
	if l == 0 {
		return 0, ErrEmpty
	}
	if l > 11 || (l == 11 && bytes[s[0]] >= 16) {
		return 0, ErrOverflow
	}
	return parse(s)
}

func FormatUint64(n uint64) string {
	buf := make([]byte, 11, 11)
	i := 10
	for n > 63 {
		buf[i] = Digits[n&63]
		i--
		n >>= 6
	}
	buf[i] = Digits[n]
	return string(buf[i-1:])
}

func FormatUint32(n uint32) string {
	return FormatUint64(uint64(n))
}
