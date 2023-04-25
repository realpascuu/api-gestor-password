package user

type User struct {
	ID       uint   `json:"id, omitempty"`
	email    string `json:"email, omitempty"`
	password string `json:"-"`
	salt     string `json:"salt, omitempty"`
	token    string `json:"-"`
}

func (u *User) HashPassword() error {
	// TODO: hash password in pbkdf2
	return nil
}
