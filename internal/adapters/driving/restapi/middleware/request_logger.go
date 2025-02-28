package middleware

import (
	"bytes"
	"io"

	"github.com/andriykusevol/aktemplategorm/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RequestLogger ...
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logger.Logger()
		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			l.Error(err.Error())
			c.Next()
		}

		l.Info(
			"Incoming request",
			zap.String("Protocol", c.Request.Proto),
			zap.String("Method", c.Request.Method),
			zap.String("URI", c.Request.RequestURI),
			zap.String("RequestBody", string(requestBody)),
		)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		c.Next()
	}
}
