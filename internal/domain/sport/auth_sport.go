package sport

import (
	"context"
	authEntity "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
)

type AuthRepository interface {
	GenerateToken(ctx context.Context, userID string) (*authEntity.Auth, error)
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	FundUserByUserName(ctx context.Context, userName string) (authEntity.User, error)
}
