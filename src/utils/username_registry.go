package utils

// UserNameRegistry holds all used usernames
type UserNameRegistry struct {
	userNames map[uint32]string
}

// NewUserNameRegistry returns a new registry
func NewUserNameRegistry() *UserNameRegistry {
	return &UserNameRegistry{
		userNames: make(map[uint32]string),
	}
}

// AddUserName adds username with associated unique ID
func (r *UserNameRegistry) AddUserName(id uint32, userName string) {
	r.userNames[id] = userName
}

// GetUserName returns the username associated with given ID
func (r *UserNameRegistry) GetUserName(id uint32) string {
	return r.userNames[id]
}
