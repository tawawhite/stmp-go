package stmp

import (
	"bytes"
	"errors"
)

var escapeCaches = [256]string{':': "%3A", '\n': "%0A", ';': "%3B", '%': "%25"}

func escape(value string) []byte {
	count := 0
	for _, c := range value {
		switch c {
		case ':', '\n', ';', '%':
			count++
		}
	}
	if count == 0 {
		return []byte(value)
	}
	buf := make([]byte, len(value)+count*2)
	j := 0
	k := 0
	for i, c := range value {
		switch c {
		case ':', '\n', ';', '%':
			copy(buf[k:], value[j:i])
			k += i - j + 3
			j = i + 1
			copy(buf[k-3:k], escapeCaches[c])
		}
	}
	copy(buf[k:], value[j:])
	return buf
}

var unescapeCaches = map[string]byte{"3A": ':', "0A": '\n', "3B": ';', "25": '%'}

func unescape(value string) string {
	max := len(value)
	buf := make([]byte, max, max)
	j := 0
	k := 0
	safe := max - 2
	for i := 0; i < safe; i++ {
		if value[i] == '%' {
			switch value[i+1 : i+3] {
			case "3A", "0A", "3B", "25":
				copy(buf[k:], value[j:i])
				k += i - j
				buf[k] = unescapeCaches[value[i+1:i+3]]
				k++
				i += 2
				j = i + 1
			}
		}
	}
	copy(buf[k:], value[j:])
	return string(buf[:k+max-j])
}

type Header map[string][]string

func (h Header) Get(key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (h Header) GetAll(key string) []string {
	return h[key]
}

func (h Header) Has(key string) bool {
	return h[key] != nil
}

func (h Header) Add(key string, value string) {
	v := h[key]
	if v == nil {
		h[key] = []string{value}
	} else {
		h[key] = append(h[key], value)
	}
}

func (h Header) Set(key string, value string) {
	h[key] = []string{value}
}

func (h Header) Del(key string) {
	delete(h, key)
}

func (h Header) Clear() {
	for k := range h {
		delete(h, k)
	}
}

func (h Header) Marshal() []byte {
	if len(h) == 0 {
		return nil
	}
	buf := make([]byte, 0, len(h)*25)
	for k, vs := range h {
		bk := escape(k)
		for _, v := range vs {
			buf = append(buf, bk...)
			buf = append(buf, ':')
			buf = append(buf, escape(v)...)
			buf = append(buf, ';')
		}
	}
	return buf[:len(buf)-1]
}

func (h Header) Unmarshal(data []byte) error {
	var sep int
	var end int
	var key string
	var value string
	for len(data) > 0 {
		sep = bytes.IndexByte(data, ':')
		if sep < 0 {
			return errors.New("miss ':' in header")
		}
		key = unescape(string(data[:sep]))
		data = data[sep+1:]
		end = bytes.IndexByte(data, ';')
		if end == -1 {
			value = unescape(string(data))
			data = nil
		} else {
			value = unescape(string(data[:end]))
			data = data[end+1:]
		}
		if len(h[key]) > 0 {
			h[key] = append(h[key], value)
		} else {
			h[key] = []string{value}
		}
	}
	return nil
}

func NewHeader() Header {
	return map[string][]string{}
}
