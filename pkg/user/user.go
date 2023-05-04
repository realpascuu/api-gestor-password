package user

type User struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	Salt     string `json:"salt,omitempty"`
}

func (u *User) HashPassword() error {
	// TODO: hash password in pbkdf2
	return nil
}
