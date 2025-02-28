package response

type FeeReponse struct {
	Amount      float64 `json:"amount" binding:"required"`
	FeeAmount   float64 `json:"fee_amount" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
	PspCode     string  `json:"psp_code" binding:"required"`
}

type BulkFeeResponse struct {
	Results []FeeReponse `json:"results" binding:"required"`
}

func NewFeeResponse(amount, fees, totalAmount float64, pspCode string) FeeReponse {

	return FeeReponse{
		Amount:      amount,
		FeeAmount:   fees,
		TotalAmount: totalAmount,
		PspCode:     pspCode,
	}
}
