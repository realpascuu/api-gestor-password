package data

import (
	"context"
	"gestorpasswordapi/pkg/user"
)

type UserRepository struct {
	Data *Data
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]user.User, error) {
	q := `SELECT id, email, password, salt from users`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.Email, &u.Password, &u.Salt)
		users = append(users, u)
	}

	return users, nil
}

func (ur *UserRepository) GetAllNotSensitiveInfo(ctx context.Context) ([]user.User, error) {
	q := `SELECT id, email from users`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.Email)
		users = append(users, u)
	}

	return users, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, id uint) (user.User, error) {
	q := `SELECT id, email from users WHERE id = $1`

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Email)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (user.User, error) {
	q := `SELECT id, email, password, salt from users WHERE email = $1`

	row := ur.Data.DB.QueryRowContext(ctx, q, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Salt)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (ur *UserRepository) Create(ctx context.Context, u *user.User) error {
	q := `INSERT INTO users (email, password, salt) VALUES ($1, $2, $3) RETURNING id`

	row := ur.Data.DB.QueryRowContext(ctx, q, u.Email, u.Password, u.Salt)

	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, u user.User) error {
	q := `UPDATE users SET email = $1, password = $2, salt = $3 WHERE id = $4`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.Email, u.Password, u.Salt, u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM users WHERE id = $1`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
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
