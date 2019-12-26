// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 16:03:13
package stmp

import "strings"

type Headers map[string][]string

func (m Headers) Get(key string) string {
	if v := m[strings.ToLower(key)]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (m Headers) GetAll(key string) []string {
	return m[strings.ToLower(key)]
}

func (m Headers) Has(key string) bool {
	return m[strings.ToLower(key)] != nil
}

func (m Headers) Add(key string, value ...string) {
	key = strings.ToLower(key)
	v := m[key]
	if v == nil {
		m[key] = value
	} else {
		m[key] = append(m[key], value...)
	}
}

func (m Headers) Set(key string, value ...string) {
	m[strings.ToLower(key)] = value
}

func (m Headers) Del(key string) {
	delete(m, strings.ToLower(key))
}

func NewHeaders() Headers {
	return map[string][]string{}
}
