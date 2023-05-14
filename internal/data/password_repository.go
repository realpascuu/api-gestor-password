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

func (pr *PasswordsRepository) GetAll(ctx context.Context, user_id uint) ([]passwords.Passwords, error) {
	q := `SELECT id, content, user_id, updated_at FROM passwords WHERE user_id = $1`

	rows, err := pr.Data.DB.QueryContext(ctx, q, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var passwordList []passwords.Passwords

	for rows.Next() {
		var p passwords.Passwords
		rows.Scan(&p.ID, &p.Content, &p.UserID, &p.UpdatedAt)
		passwordList = append(passwordList, p)
	}

	return passwordList, nil
}

func (pr *PasswordsRepository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM passwords WHERE id = $1`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
