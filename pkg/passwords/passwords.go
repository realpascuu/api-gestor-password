package passwords

import "time"

type Passwords struct {
	ID        string     `json:"id,omitempty"`
	Content   string     `json:"content"`
	UserID    uint       `json:"user_id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
