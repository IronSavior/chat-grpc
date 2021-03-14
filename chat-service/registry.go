package chat

import (
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// registry tracks the set of running chats for the service
type registry struct {
	db map[string]*Chat

	mtx sync.RWMutex
}

// Add ...
func (r *registry) Add(name string, conf *Chat) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.db[name]; ok {
		return status.Errorf(codes.AlreadyExists, "chat name %q already exists", name)
	}

	r.db[name] = conf
	return nil
}

// Remove ...
func (r *registry) Remove(name string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.db[name]; !ok {
		return status.Errorf(codes.NotFound, "no such chat: %q", name)
	}

	delete(r.db, name)
	return nil
}

// Get ...
func (r *registry) Get(name string) (*Chat, bool) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	conf, ok := r.db[name]
	return conf, ok
}

// Chats returns the list of registered chats
func (r *registry) Chats() []*Chat {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	var chats []*Chat
	for _, chat := range r.db {
		chats = append(chats, chat)
	}

	return chats
}
