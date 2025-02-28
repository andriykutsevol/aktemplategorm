// Package middleware ...
package middleware

import (
	//"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CORSMiddleware function ...
func CORSMiddleware() gin.HandlerFunc {

	//TODO: read from cfg
	return cors.New(cors.Config{
		AllowAllOrigins:  true,                                                         // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // Allow all methods
		AllowHeaders:     []string{"*"},                                                // Allow all headers
		ExposeHeaders:    []string{"*"},                                                // Expose all headers
		AllowCredentials: true,                                                         // Allow credentials
		MaxAge:           12 * time.Hour,                                               // Preflight request cache duration
	})

	// return func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Methods", "OPTIONS, POST, PUT, PATCH, GET, DELETE")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-MRA-SOURCE")

	// 	// // if method is OPTIONS then I wanna stop and return with OK status
	// 	// if c.Request.Method == http.MethodOptions {
	// 	// 	c.String(http.StatusOK, "")
	// 	// }

	// 	// Execute pending handlers in the chain
	// 	c.Next()
	// }
}