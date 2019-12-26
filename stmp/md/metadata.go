// Copyright 2019 yangjunbao <yangjunbao@shimo.im>. All rights reserved.
// Since 2019-12-23 16:03:13
package md

type Metadata map[string][]string

func (m Metadata) Get(key string) string {
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0]
	}
	return ""
}

func (m Metadata) GetAll(key string) []string {
	if v, ok := m[key]; ok && len(v) > 0 {
		return v
	}
	return nil
}

func (m Metadata) Has(key string) bool {
	if v, ok := m[key]; ok {
		return len(v) > 0
	}
	return false
}

func (m Metadata) Add(key string, value ...string) {
	if m.Has(key) {
		m[key] = append(m[key], value...)
	} else {
		m[key] = value
	}
}

func (m Metadata) Set(key string, value ...string) {
	m[key] = value
}

func (m Metadata) Del(key string) {
	if m.Has(key) {
		delete(m, key)
	}
}

func New() Metadata {
	return map[string][]string{}
}
