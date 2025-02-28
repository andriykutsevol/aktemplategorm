package request

import (
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/auditdbrecord"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"

	"github.com/google/uuid"
)

// PSP struct for PSP
type PSPObject struct {
	// Country code ISO 3166-1 ALPHA-2
	PspCountryCode *string `json:"PspCountryCode,omitempty"`
	// PSP code in a specific format (Country_code - PSP_shortcode)
	PspCode string `json:"PspCode" binding:"required"`
	// Short human friendly name of the PSPShort human friendly name of the PSP
	PspShortName *string `json:"PspShortName,omitempty"`
}

type PSPQuery struct {
	PaginationParam queryfilter.PaginationQuery // c.ShouldBindQuery will find annotations and bind them.
	PspCode         *string                     `form:"PspCode"`
	PspCountryCode  *string                     `form:"PspCountryCode"`
	IDs             *[]string                   `form:"IDs"`
	Sort            *[]string                   `form:"Sort"`
}

func (qp *PSPQuery) ToDomain_PSPQuery() *psp.Query {

	domainQP := psp.NewQuery(qp.PaginationParam)

	if qp.Sort != nil {
		of, _ := queryfilter.OrderFieldsFromStrings(*qp.Sort)
		domainQP.OrderFields = of
	}

	domainQP.IDs = qp.IDs
	domainQP.PspCode = qp.PspCode
	domainQP.PspCountryCode = qp.PspCountryCode

	return domainQP
}

func (p *PSPObject) ToDomain_PSPObject(userID string) *psp.PSP {

	//TODO: Validate request inputs.

	//TODO: Handle error
	createdBy, _ := uuid.Parse(userID)

	a := auditdbrecord.NewAuditDBRecord(createdBy)

	psp := psp.New(p.PspCode, a)

	if p.PspCountryCode != nil {
		psp.PspCountryCode = p.PspCountryCode
	}

	if p.PspShortName != nil {
		psp.PspShortName = p.PspShortName
	}
	return &psp
}
