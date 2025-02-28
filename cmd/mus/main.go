// Package main is the entry point for the application.
package main

import (
	"github.com/andriykusevol/aktemplategorm/internal/application"
	"os"
)

func main() {
	application.Run(
		os.Getenv("APP_MODE"),
		os.Getenv("APP_PORT"),
		os.Getenv("DATABASE_DSN"),
		os.Getenv("MYSQL_MAX_OPENCONNS"),
		os.Getenv("MYSQL_MAX_IDLECONS"),
		os.Getenv("COMPONENT"),
		os.Getenv("API_VERSION"),
		os.Getenv("ENV"),
	)
}
