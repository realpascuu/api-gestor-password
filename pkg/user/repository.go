package user

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetAllNotSensitiveInfo(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}
