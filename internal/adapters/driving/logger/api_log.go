package logger

import "go.uber.org/zap"

type ApiLogDataFieldName string

const (
	TIMESTAMP        ApiLogDataFieldName = "timestamp"
	LOG_LEVELS       ApiLogDataFieldName = "log_levels"
	CORRELATION_ID   ApiLogDataFieldName = "cid"
	EVENT_OWNER      ApiLogDataFieldName = "event_owner"
	PROCESS          ApiLogDataFieldName = "process"
	SUB_PROCESS      ApiLogDataFieldName = "sub_process"
	PROCESSED_OBJECT ApiLogDataFieldName = "processed_object"
	STATUS           ApiLogDataFieldName = "status"
	RESULT           ApiLogDataFieldName = "result"
	ERROR_MESSAGE    ApiLogDataFieldName = "error_message"
	ERROR_CODE       ApiLogDataFieldName = "error_code"
	API_ENDPOINT     ApiLogDataFieldName = "api_endpoint"
	API_METHOD       ApiLogDataFieldName = "api_method"
	CLIENT_IP        ApiLogDataFieldName = "client_ip"
	PAYLOAD          ApiLogDataFieldName = "payload"
	HOSTNAME         ApiLogDataFieldName = "hostname"

	COMPONENT   ApiLogDataFieldName = "component"
	ENVIRONMENT ApiLogDataFieldName = "env"
	BUSINESS_ID ApiLogDataFieldName = "business_id"
	API_VERSION ApiLogDataFieldName = "api_version"
)

type ApiLoggerIface interface {
	CreateApiLogData(fields ...ApiLogDataField) map[ApiLogDataFieldName]interface{}
	CreateApiLogPayload(fields ...ApiLogDataPayloadField) map[string]interface{}
}

type ApiLogDataField struct {
	Key   ApiLogDataFieldName
	Value interface{}
}

type ApiLogDataPayloadField struct {
	Key   string
	Value interface{}
}

type ApiLogGenerator struct {
	Component   string
	Environment string
	ApiVersion  string
	Logger      *zap.Logger
}

func NewApiLogGenerator(component, environement, apiVersion string, logger *zap.Logger) ApiLoggerIface {
	return &ApiLogGenerator{
		Component:   component,
		Environment: environement,
		ApiVersion:  apiVersion,
		Logger:      logger,
	}
}

func (apiLogData *ApiLogGenerator) CreateApiLogData(fields ...ApiLogDataField) map[ApiLogDataFieldName]interface{} {

	result := make(map[ApiLogDataFieldName]interface{})

	//load default fields from .env
	result[COMPONENT] = apiLogData.Component
	result[ENVIRONMENT] = apiLogData.Environment
	result[API_VERSION] = apiLogData.ApiVersion

	//add the provided fields in the log data
	for _, field := range fields {
		result[field.Key] = field.Value
	}

	return result
}

func (apiLogData *ApiLogGenerator) CreateApiLogPayload(fields ...ApiLogDataPayloadField) map[string]interface{} {

	result := make(map[string]interface{})
	//add the provided fields in the log data
	for _, field := range fields {
		result[field.Key] = field.Value
	}

	return result
}
