package app

import (
	"context"
	"errors"
	"github.com/andriykusevol/aktemplategorm/internal/application/pport"
	authEntity "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
)

// type Auth interface {
// 	Verify(ctx context.Context, userName, password string) (*auth_domain.User, error)
// 	GenerateToken(ctx context.Context, userID string) (*auth_domain.Auth, error)
// }

type authApp struct {
	authRepo sport.AuthRepository
}

func NewAuthApp(authRepo sport.AuthRepository) pport.AuthApp {
	return &authApp{
		authRepo: authRepo,
	}
}

func (app authApp) Verify(ctx context.Context, userName, password string) (*authEntity.User, error) {

	user, err := app.authRepo.FundUserByUserName(ctx, userName)
	if err != nil {
		return nil, err //TODO: Domain Error: errors.ErrInvalidUserName
	}

	if user.Password != password {
		return nil, errors.New("Invalid password") //TODO: Eomain Error: errors.ErrInvalidPassword
	}

	return &user, nil
}

func (app authApp) GenerateToken(ctx context.Context, userUUID string) (*authEntity.Auth, error) {

	auth, err := app.authRepo.GenerateToken(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
