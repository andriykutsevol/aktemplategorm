package pport

import (
	"context"
	auth_domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
)

type AuthApp interface {
	Verify(ctx context.Context, userName, password string) (*auth_domain.User, error)
	GenerateToken(ctx context.Context, userID string) (*auth_domain.Auth, error)
}
