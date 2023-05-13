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

func (pr *PasswordsRepository) GetOne(ctx context.Context, id string) (passwords.Passwords, error) {
	q := `SELECT id, content, user_id, updated_at FROM passwords WHERE id = $1`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)
	var p passwords.Passwords
	err := row.Scan(&p.ID, &p.Content, &p.UserID, &p.UpdatedAt)
	if err != nil {
		return passwords.Passwords{}, err
	}

	return p, nil
}
