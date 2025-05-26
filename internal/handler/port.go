package handler

import (
	"container-manager/internal/schema"
	"context"
)

type UserService interface {
}

type AuthService interface {
	Login(ctx context.Context, req *schema.User) (string, error)
	NewUser(ctx context.Context, req *schema.User) error
}
