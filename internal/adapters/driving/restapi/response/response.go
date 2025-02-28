// Package response ...
package response

// CustomErrMsg ...
type CustomErrMsg struct {
	Language string `json:"Language"`
	Content  string `json:"Content"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Raw               int          `json:"Raw"`
	ProviderErrorCode string       `json:"ProviderErrorCode"`
	MappedErrorCode   string       `json:"MappedErrorCode"`
	DevMsg            string       `json:"DevMsg"`
	Serverity         string       `json:"Severity"`
	CustomMsg         CustomErrMsg `json:"CustomMsg"`
}
