package request

type FeeRequest struct {
	PspID   string `json:"PspID"`
	PspCode string `json:"PspCode"`
	Amount  string `json:"Amount"`
}

type BulkFeeRequest []FeeRequest
