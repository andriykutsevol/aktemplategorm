package middleware

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"

	apierror "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/errors"
	auth_domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/auth"
)

// https://community.postman.com/t/fixing-or-ignoring-response-header-issues/53550

// The Date header is typically set to the current server time in the RFC 1123 format,
// which is the standard format for HTTP headers.
func DateHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the current server time in RFC1123 format
		// currentTime := time.Now().UTC().Format(time.RFC1123)

		manualTime := time.Date(
			2024,          // Year
			time.December, // Month
			3,             // Day
			14,            // Hour
			33,            // Minute
			0,             // Second
			0,             // Nanosecond
			time.UTC,      // Timezone (UTC in this case)
		)
		formattedTime := manualTime.Format(time.RFC1123)

		// Set the "Date" header in the response
		c.Writer.Header().Set("Date", formattedTime)

		// Proceed with the request
		c.Next()
	}
}

func RequestHeadersLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=== Request Headers ===")
		for key, values := range c.Request.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
		fmt.Println("=======================")

		// Proceed to the next handler
		c.Next()
	}
}

func ResponseHeadersLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Proceed with the request
		c.Next()

		// Log response headers after the request is handled
		fmt.Println("=== Response Headers ===")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
		fmt.Println("=======================")
	}
}

// CustomResponseWriter to capture the response headers and body
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(data []byte) (int, error) {
	// Write the body to the buffer (for logging)
	w.body.Write(data)
	// Write the body to the actual ResponseWriter
	return w.ResponseWriter.Write(data)
}

func LogResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a custom writer
		writer := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		// Replace the default writer with the custom writer
		c.Writer = writer

		// Process the request
		c.Next()

		// Log the response headers
		fmt.Fprintln(os.Stdout, "=== Response Headers ===")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				fmt.Fprintf(os.Stdout, "%s: %s\n", key, value)
			}
		}
		fmt.Fprintln(os.Stdout, "=======================")

		// Log the response body
		fmt.Fprintln(os.Stdout, "\n=== Response Body ===")
		io.Copy(os.Stdout, writer.body)
		fmt.Fprintln(os.Stdout) // Add a newline for better formatting
		fmt.Fprintln(os.Stdout, "=======================")

	}
}

func XRateLimitLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-RateLimit-Limit", "100500")
		// Proceed with the request
		c.Next()
	}
}

func NewRequestCID(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, auth_domain.CIDCtx{}, cid)
}

func wrapRequestCIDContext(c *gin.Context, cid string) {
	c.Set("X-Correlation-Id", cid)
	// We store this not just in gin context, but in the "context" itself also.
	ctx := NewUserID(c.Request.Context(), cid) // To access it: ctxUid := ctx.Value(auth_domain.UserIDCtx{})
	c.Request = c.Request.WithContext(ctx)
}

func XCorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//--------------------------------------------------
		// Correlation ID.

		correlationID := c.GetHeader("X-Correlation-Id")
		if len(correlationID) == 0 {
			entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
			ms := ulid.Timestamp(time.Now())
			cid, err := ulid.New(ms, entropy)
			if err != nil {
				errObj := apierror.NewApiError(
					"", //TODO: strconv.Itoa(http.StatusBadRequest),
					apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
					"Unexpected server error",
					apierror.SEVERITY_LOW, err,
				)
				c.JSON(http.StatusInternalServerError, errObj)
				c.Abort()
				return
			}
			correlationID = cid.String()
		}

		wrapRequestCIDContext(c, correlationID)

		//-------------------------------------------------
		// Add x-correlation-Id to header response.
		// Header is an intelligent shortcut for c.Writer.Header().Set(key, value).
		// It writes a header in the response.
		// If value == "", this method removes the header `c.Writer.Header().Del(key)`
		c.Header("X-Correlation-Id", correlationID)

		c.Next()
	}
}

type responseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	// Write data to the buffer
	rw.buffer.Write(data)
	// Do not write to the actual ResponseWriter yet
	return len(data), nil
}

// func ContentLengthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Capture the response using a custom ResponseWriter
// 		buffer := &bytes.Buffer{}
// 		writer := &responseWriter{
// 			ResponseWriter: c.Writer,
// 			buffer:         buffer,
// 		}
// 		c.Writer = writer

// 		// Process the request
// 		c.Next()

// 		// Calculate Content-Length and set the header
// 		contentLength := buffer.Len()
// 		c.Header("Content-Length", strconv.Itoa(contentLength))

// 		// Write the buffered response to the original writer
// 		_, _ = c.Writer.Write(buffer.Bytes())
// 	}
// }

func ContentLengthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		teeReader := io.TeeReader(c.Request.Body, &buf)
		c.Request.Body = io.NopCloser(teeReader)

		c.Next()

		contentLength := strconv.Itoa(len(buf.Bytes()))
		c.Header("Content-Length", contentLength)
	}
}

func InternalServerErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		errObj := apierror.NewApiError(
			"", //TODO: strconv.Itoa(http.StatusBadRequest),
			apierror.ERROR_CODE_UNEXPECTED_SERVER_ERROR,
			"Unexpected server error",
			apierror.SEVERITY_LOW, errors.New(""),
		)
		c.JSON(http.StatusInternalServerError, errObj)
		c.Abort()
		return
	}
}
