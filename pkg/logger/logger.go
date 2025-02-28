// Package logger ...
package logger

import (
	"go.uber.org/zap"
	"sync"
)

var logger *zap.Logger
var once sync.Once

// Logger - returns pointer to global logger *zap.Logger
func Logger() *zap.Logger {
	once.Do(func() {
		// Initialize the logger only once
		var err error
		logger, err = zap.NewProduction() // or zap.NewDevelopment()
		if err != nil {
			panic(err) // handle error appropriately
		}
	})
	return logger
}
