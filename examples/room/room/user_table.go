// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-02 14:26:11
package room

import (
	"github.com/acrazing/stmp-go/examples/room/room_proto"
	"sync"
)

type UserTable struct {
	mu   sync.RWMutex
	rows map[string]*room_proto.UserModel
}

func (t *UserTable) List(limit int64, offset int64) (total int64, list []*room_proto.UserModel) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	total = int64(len(t.rows))
	limit = Min(total-offset, limit)
	if limit < 1 {
		return
	}
	list = make([]*room_proto.UserModel, limit, limit)
	i := int64(0)
	for _, row := range t.rows {
		if offset == 0 {
			list[i] = row
			i += 1
			if i == limit {
				break
			}
		} else {
			offset -= 1
		}
	}
	return
}

func (t *UserTable) Size() int64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return int64(len(t.rows))
}

func (t *UserTable) Push(users []*room_proto.UserModel) {
	t.mu.Lock()
	for _, u := range users {
		t.rows[u.Name] = u
	}
	t.mu.Unlock()
}

func (t *UserTable) Put(user *room_proto.UserModel) {
	t.mu.Lock()
	t.rows[user.Name] = user
	t.mu.Unlock()
}

func (t *UserTable) Del(name string) {
	t.mu.Lock()
	delete(t.rows, name)
	t.mu.Unlock()
}

func NewUserTable() *UserTable {
	return &UserTable{
		rows: map[string]*room_proto.UserModel{},
	}
}
