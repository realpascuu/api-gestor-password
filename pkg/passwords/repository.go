package passwords

import "context"

type Repository interface {
	/* GetAll(ctx context.Context) ([]Passwords, error)
	GetOne(ctx context.Context, id string) (Passwords, error)
	GetByUser(ctx context.Context, userID uint) (Passwords, error) */
	Create(ctx context.Context, post *Passwords) error
	/* Update(ctx context.Context, id string, post Passwords) error
	Delete(ctx context.Context, id string) error */
}
