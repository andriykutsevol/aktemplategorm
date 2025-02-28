package money

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/currency"
	"github.com/govalues/decimal"
)

// We make this structure immutable
type Money struct {
	amount decimal.Decimal
	code   string
	number int16
}

func NewMoney(amount decimal.Decimal, currency currency.Currency) Money {
	return Money{amount: amount, code: currency.Code()}
}

func (m Money) Amount() decimal.Decimal {
	return m.amount
}

func (m Money) Code() string {
	return m.code
}
