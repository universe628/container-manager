package authservice

import (
	"container-manager/internal/schema"
	"context"
)

type UserRepo interface {
	CreateUser(context.Context, *schema.User) (int, error)
	GetUserByUserName(context.Context, string) (*schema.User, error)
}

type TokenProvider interface {
	GenerateToken(int32) (string, error)
}
