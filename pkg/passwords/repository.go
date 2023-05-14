package passwords

import "context"

type Repository interface {
	GetOne(ctx context.Context, id string) (Passwords, error)
	GetAll(ctx context.Context, user_id uint) ([]Passwords, error)
	Create(ctx context.Context, post *Passwords) error
	Delete(ctx context.Context, id string) error
	/* Update(ctx context.Context, id string, post Passwords) error */
}
