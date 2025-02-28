package errors

// The custom mapped error for the corporate service
// They are custom to the smartcash corporate payment service, and allow to provide more context about a given error (exception)
// check the documentation here : https://Some.atlassian.net/wiki/spaces/SCG/pages/3555983361/CPS+-+Corporate+Pay+Service+API+Error+Codes

type ApiErrorCode string

const (
	ERROR_CODE_MISSING_REQUIRED_PARAMETER         ApiErrorCode = "MUS-4001"
	ERROR_CODE_INVALID_DATA_FORMAT                ApiErrorCode = "MUS-4002"
	ERROR_CODE_INVALID_REQUEST_METHOD             ApiErrorCode = "MUS-4003"
	ERROR_CODE_REQUEST_BODY_TOO_LARGE             ApiErrorCode = "MUS-4004"
	ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS ApiErrorCode = "MUS-4011"
	ERROR_CODE_INSUFFICIENT_PRIVILEGES            ApiErrorCode = "MUS-4012"
	ERROR_CODE_AUTHENTICATION_REQUIRED            ApiErrorCode = "MUS-4013"
	ERROR_CODE_INSUFFICIENT_PERMISSIONS           ApiErrorCode = "MUS-4031"
	ERROR_CODE_RESOURCE_OWNERSHIP_VIOLATION       ApiErrorCode = "MUS-4032"
	ERROR_CODE_IP_BLOCKED                         ApiErrorCode = "MUS-4033"
	ERROR_CODE_ENDPOINT_NOT_FOUND                 ApiErrorCode = "MUS-4041"
	ERROR_CODE_RESOURCE_NOT_FOUND                 ApiErrorCode = "MUS-4042"
	ERROR_CODE_DUPLICATE_RESOURCE                 ApiErrorCode = "MUS-4091"
	ERROR_CODE_RESOURCE_ALREADY_PROCESSED         ApiErrorCode = "MUS-4092"
	ERROR_CODE_INVALID_DATA_FORMAT_4221           ApiErrorCode = "MUS-4221"
	ERROR_CODE_DATA_VALIDATION_FAILURE            ApiErrorCode = "MUS-4222"
	ERROR_CODE_UNEXPECTED_SERVER_ERROR            ApiErrorCode = "MUS-5001"
	ERROR_CODE_DATABASE_CONNECTION_FAILURE        ApiErrorCode = "MUS-5002"
)

type ApiErrorSeverity string

const (
	SEVERITY_NO     ApiErrorSeverity = "NO"
	SEVERITY_LOW    ApiErrorSeverity = "LOW"
	SEVERITY_MEDIUM ApiErrorSeverity = "MEDIUM"
	SEVERITY_HIGH   ApiErrorSeverity = "HIGH"
)

type TranslatedMessage struct {
	Language string
	Content  string
}

// check the documentation to find more on error message and translations
// link of the documentation here : https://Some.atlassian.net/wiki/spaces/SCG/pages/3555983361/CPS+-+Corporate+Pay+Service+API+Error+Codes
var translatedErrorMessages = map[ApiErrorCode][]TranslatedMessage{
	ERROR_CODE_MISSING_REQUIRED_PARAMETER: {
		{Language: "en", Content: "Missing required parameter"},
		{Language: "fr", Content: "Paramètre requis manquant"},
	},
	ERROR_CODE_INVALID_DATA_FORMAT: {
		{Language: "en", Content: "Invalid data format"},
		{Language: "fr", Content: "Format de données invalide"},
	},
	ERROR_CODE_INVALID_REQUEST_METHOD: {
		{Language: "en", Content: "Invalid request method"},
		{Language: "fr", Content: "Méthode de requête invalide"},
	},
	ERROR_CODE_REQUEST_BODY_TOO_LARGE: {
		{Language: "en", Content: "Request body too large"},
		{Language: "fr", Content: "Corps de la requête trop volumineux"},
	},
	ERROR_CODE_INVALID_AUTHENTICATION_CREDENTIALS: {
		{Language: "en", Content: "Invalid authentication credentials"},
		{Language: "fr", Content: "Identifiants d'authentification invalides"},
	},
	ERROR_CODE_INSUFFICIENT_PRIVILEGES: {
		{Language: "en", Content: "Insufficient privileges"},
		{Language: "fr", Content: "Privilèges insuffisants"},
	},
	ERROR_CODE_AUTHENTICATION_REQUIRED: {
		{Language: "en", Content: "Authentication required"},
		{Language: "fr", Content: "Authentification requise"},
	},
	ERROR_CODE_INSUFFICIENT_PERMISSIONS: {
		{Language: "en", Content: "Insufficient permissions"},
		{Language: "fr", Content: "Permissions insuffisantes"},
	},
	ERROR_CODE_RESOURCE_OWNERSHIP_VIOLATION: {
		{Language: "en", Content: "Resource ownership violation"},
		{Language: "fr", Content: "Violation de la propriété des ressources"},
	},
	ERROR_CODE_IP_BLOCKED: {
		{Language: "en", Content: "IP blocked"},
		{Language: "fr", Content: "IP bloquée"},
	},
	ERROR_CODE_ENDPOINT_NOT_FOUND: {
		{Language: "en", Content: "Endpoint not found"},
		{Language: "fr", Content: "Point de terminaison introuvable"},
	},
	ERROR_CODE_RESOURCE_NOT_FOUND: {
		{Language: "en", Content: "Resource not found"},
		{Language: "fr", Content: "Ressource introuvable"},
	},
	ERROR_CODE_DUPLICATE_RESOURCE: {
		{Language: "en", Content: "Duplicate resource"},
		{Language: "fr", Content: "Ressource dupliquée"},
	},
	ERROR_CODE_RESOURCE_ALREADY_PROCESSED: {
		{Language: "en", Content: "Resource already processed"},
		{Language: "fr", Content: "Ressource déjà traitée"},
	},
	ERROR_CODE_INVALID_DATA_FORMAT_4221: {
		{Language: "en", Content: "Invalid data format"},
		{Language: "fr", Content: "Format de données invalide"},
	},
	ERROR_CODE_DATA_VALIDATION_FAILURE: {
		{Language: "en", Content: "Data validation failure"},
		{Language: "fr", Content: "Échec de la validation des données"},
	},
	ERROR_CODE_UNEXPECTED_SERVER_ERROR: {
		{Language: "en", Content: "Unexpected server error"},
		{Language: "fr", Content: "Erreur serveur inattendue"},
	},
	ERROR_CODE_DATABASE_CONNECTION_FAILURE: {
		{Language: "en", Content: "Database connection failure"},
		{Language: "fr", Content: "Échec de la connexion à la base de données"},
	},
}

type ApiError struct {
	Message           string               `json:"raw"`                         //Error message
	ProviderErrorCode *string              `json:"providerErrorCode,omitempty"` //Provider error code : an error describing an exception that occured in a third party service (one of the external service with wich we interact for some reason )
	MappedErrorCode   *ApiErrorCode        `json:"mappedErrorCode,omitempty"`   //custom error code belonging to smartcash corporate servie check the comments at the bigening of the file.
	DevMsg            *string              `json:"devMsg,omitempty"`            //Error mesage for developer
	Severity          *ApiErrorSeverity    `json:"severity,omitempty"`          //Error severity from "NO" to "HIGHT"
	CustomerMsg       *[]TranslatedMessage `json:"customerMsg,omitempty"`       //User friendly error messages to display on the front-end (with or without their transalations)
	PreviousError     *error               `json:"-"`                           //The previous error (mostly the one that led us to create this custom error) used for the error chaining
}

func NewApiError(
	providerErrorCode string,
	mappedErrorCode ApiErrorCode,
	devMsg string,
	severity ApiErrorSeverity,
	previousError error) ApiError {

	customerMsg := translatedErrorMessages[mappedErrorCode]

	return ApiError{
		Message:           previousError.Error(),
		ProviderErrorCode: &providerErrorCode,
		MappedErrorCode:   &mappedErrorCode,
		DevMsg:            &devMsg,
		Severity:          &severity,
		CustomerMsg:       &customerMsg,
		PreviousError:     &previousError,
	}
}
