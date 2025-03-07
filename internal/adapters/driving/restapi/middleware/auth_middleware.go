package middleware

import (
	"context"
	"fmt"
	apierror "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/errors"
	auth_domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// The reason for using an empty struct like this
// is to ensure that the key used to store the user ID in the context is unique.
// This pattern prevents accidental conflicts with other context values,
// even if other parts of the code store values in the context.
func NewUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, auth_domain.UserIDCtx{}, userID)
}

func wrapUserAuthContext(c *gin.Context, userID string) {
	c.Set("UserID", userID)
	// We store this not just in gin context, but in the "context" itself also.
	ctx := NewUserID(c.Request.Context(), userID) // To access it: ctxUid := ctx.Value(auth_domain.UserIDCtx{})
	c.Request = c.Request.WithContext(ctx)
}

func AuthMiddleware(r sport.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		//--------------------------------------------------
		// Parse User
		var token string
		auth := c.GetHeader("x-api-key")

		prefix := "Bearer "
		if auth != "" && strings.HasPrefix(auth, prefix) {
			token = auth[len(prefix):]
		}

		//TODO: Handle error.
		userUUID, err := r.ParseUserID(c.Request.Context(), token)

		if err != nil {
			errObj := apierror.NewApiError(
				"", //TODO: strconv.Itoa(http.StatusBadRequest),
				apierror.ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS,
				"http.StatusUnauthorized: Cannot parse token.",
				apierror.SEVERITY_LOW, err,
			)
			c.JSON(http.StatusUnauthorized, errObj)
			c.Abort()
			return
		}
		wrapUserAuthContext(c, userUUID)
		c.Next()
	}
}
