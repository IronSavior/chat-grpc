package chat

// User represents a connected chat user with command/event streams
// NB: User values MUST be comparable, therefore Stream interface values MUST ALSO be comparable
type User struct {
	name string
	Events
	CommandStream
}

// Name returns the name of the User
func (u User) Name() string {
	return u.name
}
