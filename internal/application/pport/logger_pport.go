package pport

import (
	loggeradapter "github.com/andriykusevol/aktemplategorm/internal/adapters/driving/logger"
)

type ApiLogger interface {
	CreateApiLogData(fields ...loggeradapter.ApiLogDataField)
	CreateApiLogPayload(fields ...loggeradapter.ApiLogDataPayloadField)
}
