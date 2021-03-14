package chat

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// roster is a set of chat members (not thread safe)
type roster struct {
	db map[User]struct{}
}

// Add a user to the roster
func (r *roster) Add(u User) error {
	if r.db == nil {
		r.db = map[User]struct{}{}
	}

	if _, ok := r.db[u]; ok {
		return status.Errorf(codes.AlreadyExists, "Already a member of this chat")
	}

	r.db[u] = struct{}{}

	return nil
}

// Remove a user from the roster
func (r *roster) Remove(u User) {
	delete(r.db, u)
}

func (r *roster) Has(u User) bool {
	_, ok := r.db[u]
	return ok
}

// Users returns the list of roster members
func (r *roster) Users() []User {
	var users []User
	for m := range r.db {
		users = append(users, m)
	}

	return users
}
