package authservice

import (
	errs "container-manager/internal/errors"
	"container-manager/internal/infra/auth"
	"container-manager/internal/schema"
	"context"
	"database/sql"
)

type authService struct {
	repo          UserRepo
	tokenProvider TokenProvider
}

func NewAuthService(repo UserRepo, tokenProvider TokenProvider) *authService {
	return &authService{
		repo:          repo,
		tokenProvider: tokenProvider,
	}
}

func (a *authService) Login(ctx context.Context, req *schema.User) (string, error) {
	user, err := a.repo.GetUserByUserName(ctx, req.UserName)
	if err == sql.ErrNoRows {
		return "", errs.ErrUserNotFound
	} else if err != nil {
		return "", err
	}

	if user == nil || user.Password != auth.PasswordHashing(req.Password) {
		return "", errs.ErrInvalidCredentials
	}
	return a.tokenProvider.GenerateToken(user.UId)
}

func (a *authService) NewUser(ctx context.Context, req *schema.User) error {
	_, err := a.repo.GetUserByUserName(ctx, req.UserName)
	if err != sql.ErrNoRows {
		return errs.ErrUserNameTaken
	} else if err != nil {
		return err
	}

	req.Password = auth.PasswordHashing(req.Password)

	_, err = a.repo.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}
