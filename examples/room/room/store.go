// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-02 14:26:11
package room

import (
	"github.com/acrazing/stmp-go/examples/room/room_proto"
	"sync"
)

type UserStore struct {
	Mu    *sync.RWMutex
	Users map[string]*room_proto.UserModel
}

func (us *UserStore) Size() int {
	us.Mu.RLock()
	defer us.Mu.RUnlock()
	return len(us.Users)
}

func (us *UserStore) Push(users []*room_proto.UserModel) {
	us.Mu.Lock()
	for _, u := range users {
		us.Users[u.Name] = u
	}
	us.Mu.Unlock()
}

func (us *UserStore) Put(user *room_proto.UserModel) {
	us.Mu.Lock()
	us.Users[user.Name] = user
	us.Mu.Unlock()
}

func (us *UserStore) Del(name string) {
	us.Mu.Lock()
	delete(us.Users, name)
	us.Mu.Unlock()
}

func NewUserStore() *UserStore {
	return &UserStore{
		Mu:    &sync.RWMutex{},
		Users: map[string]*room_proto.UserModel{},
	}
}
