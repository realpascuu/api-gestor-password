package passwords

import "time"

type Passwords struct {
	ID        uint      `json:"id, omitempty"`
	content   string    `json:-`
	UserID    uint      `json:user_id,omitempty`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
