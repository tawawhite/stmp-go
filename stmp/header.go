package stmp

import (
	"bytes"
	"errors"
	"strings"
)

func escapeHeadKey(key string) string {
	return strings.ReplaceAll(escapeHeadValue(key), ":", "%3A")
}

func escapeHeadValue(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "%", "%25"), "\n", "%0A")
}

func unescapeHeadKey(key string) string {
	return strings.TrimSpace(strings.ReplaceAll(unescapeHeadValue(key), "%3A", ":"))
}

func unescapeHeadValue(value string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(value, "%0A", "\n"), "%25", "%"))
}

type Header map[string][]string

func (h Header) Get(key string) string {
	if v := h[strings.ToLower(key)]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (h Header) GetAll(key string) []string {
	return h[strings.ToLower(key)]
}

func (h Header) Has(key string) bool {
	return h[strings.ToLower(key)] != nil
}

func (h Header) Add(key string, value ...string) {
	key = strings.ToLower(key)
	v := h[key]
	if v == nil {
		h[key] = value
	} else {
		h[key] = append(h[key], value...)
	}
}

func (h Header) Set(key string, value ...string) {
	h[strings.ToLower(key)] = value
}

func (h Header) Del(key string) {
	delete(h, strings.ToLower(key))
}

func (h Header) Marshal() []byte {
	if len(h) == 0 {
		return nil
	}
	chunks := make([]string, 0, len(h))
	for k, vs := range h {
		for _, v := range vs {
			chunks = append(chunks, escapeHeadKey(k)+":"+escapeHeadValue(v))
		}
	}
	return []byte(strings.Join(chunks, "\n"))
}

func (h Header) Unmarshal(data []byte) error {
	var sep int
	var end int
	var key string
	var value string
	for {
		sep = bytes.IndexByte(data, ':')
		if sep < 0 {
			return errors.New("miss ':' in header")
		}
		key = unescapeHeadKey(string(data[:sep]))
		end = bytes.IndexByte(data[sep+1:], '\n')
		if end == -1 {
			value = unescapeHeadValue(string(data[sep+1:]))
		} else {
			value = unescapeHeadValue(string(data[sep+1 : end]))
		}
		if len(h[key]) > 0 {
			h[key] = append(h[key], value)
		} else {
			h[key] = []string{value}
		}
		if end == -1 || end == len(data) {
			return nil
		}
		data = data[end+1:]
	}
}

func NewHeader() Header {
	return map[string][]string{}
}
