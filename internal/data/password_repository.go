package data

import (
	"context"
	"gestorpasswordapi/pkg/passwords"
)

type PasswordsRepository struct {
	Data *Data
}

func (pr *PasswordsRepository) Create(ctx context.Context, password *passwords.Passwords) error {
	q := `INSERT INTO passwords (content, user_id) VALUES ($1, $2) RETURNING id`

	row := pr.Data.DB.QueryRowContext(ctx, q, password.Content, password.UserID)

	err := row.Scan(&password.ID)
	if err != nil {
		return err
	}

	return nil
}
