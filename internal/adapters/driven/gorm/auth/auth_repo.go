package auth

import (
	"context"
	"errors"
	authEntity "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
	"time"

	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

//var Users []auth_domain.User

const defaultKey = "defaultKey to sing JWT"

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyFunc       jwt.Keyfunc
	expired       int
	tokenType     string
}

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       172000,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyFunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, authEntity.ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

func strPtr(s string) *string {
	return &s
}

type repository struct {
	db    *gorm.DB
	opts  *options
	users []authEntity.User //TODO: mock object instead of the database.
}

func NewRepository(db *gorm.DB) sport.AuthRepository {
	o := defaultOptions

	UserUUID, _ := uuid.Parse("de95ca9d-4898-46f4-af79-b1e70326b4c1")
	AdminUUID, _ := uuid.Parse("de95ca9d-4898-46f4-af79-b1e70326b4c2")

	return &repository{
		db:   db,
		opts: &o,
		users: []authEntity.User{
			{ID: strPtr("1"), UserName: "User", Password: "1234", UUID: &UserUUID}, //TODO: Hash passwords.
			{ID: strPtr("1"), UserName: "Admin", Password: "12345", UUID: &AdminUUID},
		},
	}
}

func (r *repository) FundUserByUserName(ctx context.Context, userName string) (authEntity.User, error) {
	for _, user := range r.users {
		if userName == user.UserName {
			return user, nil
		}
	}
	return authEntity.User{}, errors.New("User not found")
}

func (r *repository) GenerateToken(ctx context.Context, userUUID string) (*authEntity.Auth, error) {

	now := time.Now()
	expiresAt := now.Add(time.Duration(r.opts.expired) * time.Second).Unix()

	token := jwt.NewWithClaims(r.opts.signingMethod, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   userUUID,
	})

	tokenString, err := token.SignedString(r.opts.signingKey)
	if err != nil {
		return nil, err
	}

	auth := &authEntity.Auth{
		ExpiresAt:   expiresAt,
		TokenType:   r.opts.tokenType,
		AccessToken: tokenString,
	}

	return auth, nil
}

func (r *repository) parseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, r.opts.keyFunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, authEntity.ErrInvalidToken
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

func (r *repository) ParseUserID(ctx context.Context, accessToken string) (string, error) {

	if accessToken == "" {
		return "", authEntity.ErrInvalidToken
	}

	claims, err := r.parseToken(accessToken)
	if err != nil {
		return "", err
	}

	return claims.Subject, nil
}
