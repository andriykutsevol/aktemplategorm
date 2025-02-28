package currency

// https://www.iban.com/currency-codes
// ISO 4217

type Currency struct {
	country string
	nameEN  string
	code    string
	number  int
}

func NewCurrency(ccode string) Currency {
	return Currency{
		code: ccode,
	}
}

func (c Currency) Code() string {
	return c.code
}

//TOD: Add setter and getters for optional fileds.
