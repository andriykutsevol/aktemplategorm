package response

import (
	domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
)

type auth struct {
	AccessToken string `json:"AccessToken" binding:"required"`
	TokenType   string `json:"TokenType" binding:"required"`
	ExpiresAt   int64  `json:"ExpiresAt" binding:"required"`
}

func FromDomain_Auth(entity domain.Auth) (auth, error) {

	response := auth{
		AccessToken: entity.AccessToken,
		TokenType:   entity.TokenType,
		ExpiresAt:   entity.ExpiresAt,
	}

	return response, nil

}
