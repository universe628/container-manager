package postgresRepo

import (
	"container-manager/internal/schema"
	authservice "container-manager/internal/service/auth"
	"context"
)

type PostgresUserRepository struct {
	q *Queries
}

func NewPostgresUserRepository(db DBTX) authservice.UserRepo {
	return &PostgresUserRepository{
		q: New(db),
	}
}
func (u PostgresUserRepository) GetUserByUserName(ctx context.Context, userName string) (*schema.User, error) {
	res, err := u.q.GetUserByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	return &schema.User{
		UId:      res.ID,
		UserName: res.Name,
		Password: res.Pwd,
	}, nil
}

func (u PostgresUserRepository) CreateUser(ctx context.Context, user *schema.User) (int, error) {

	uid, err := u.q.CreateUser(ctx, CreateUserParams{
		Name: user.UserName,
		Pwd:  user.Password,
	})
	if err != nil {
		return 0, err
	}

	return int(uid), nil
}
