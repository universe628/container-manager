package authservice

import (
	"container-manager/internal/schema"
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *schema.User) (int, error)
	GetUserByUserName(ctx context.Context, userName string) (*schema.User, error)
}

type TokenProvider interface {
	GenerateToken(userId int32) (string, error)
}
