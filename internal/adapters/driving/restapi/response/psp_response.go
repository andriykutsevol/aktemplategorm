package response

import (
	"fmt"
	domain "github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"
)

type PspConverterFromDomain interface {
	FromDomain_Psp(psp domain.PSP) (Psp, error)
	FromDomainList_PspList(domainItems []domain.PSP) (*ListPsp, error)
	FromDomain_PaginatedPspList(domainItems []domain.PSP, p queryfilter.Pagination) (*PaginatedPspList, error)
}

type Psp struct {
	ID uint `json:"ID"`
	// Country code ISO 3166-1 ALPHA-2
	PspCountryCode *string `json:"PspCountryCode,omitempty"`
	// PSP code in a specific format (Country_code - PSP_shortcode)
	PspCode string `json:"PspCode" binding:"required"`
	// Short human friendly name of the PSPShort human friendly name of the PSP
	PspShortName *string `json:"PspShortName,omitempty"`

	A AuditDBRecord `json:"AuditDBRecord" binding:"required"`

	// CreatedAt time.Time `json:"createdat" binding:"required"`
	// CreatedBy string    `json:"createdby" binding:"required"`
}

func NewPspConverterFromDomain() PspConverterFromDomain {
	return Psp{}
}

func (resp Psp) FromDomain_Psp(psp domain.PSP) (Psp, error) {
	fmt.Println(psp)
	pspResponse := Psp{
		ID:             *psp.ID, // If it will be nil - you're cannot dereference it (nil pointer dereference)
		PspCountryCode: psp.PspCountryCode,
		PspShortName:   psp.PspShortName,
		PspCode:        psp.PspCode,
		A:              FromDomain_AuditDBRecord(psp.AuditDBRecord),
	}

	return pspResponse, nil
}

func (resp Psp) FromDomainList_PspList(domainItems []domain.PSP) (*ListPsp, error) {
	//psps := []Psp{}

	listPsp := ListPsp{
		Results: []Psp{}, // Initialize with an empty slice
	}

	for _, psp := range domainItems {
		response, _ := FromDomain_Psp(psp)
		listPsp.Results = append(listPsp.Results, response)
	}

	return &listPsp, nil
}

func (resp Psp) FromDomain_PaginatedPspList(domainItems []domain.PSP, p queryfilter.Pagination) (*PaginatedPspList, error) {
	//psps := []Psp{}

	response := PaginatedPspList{
		Results:    []Psp{}, // Initialize with an empty slice
		Pagination: p,
	}

	for _, item := range domainItems {
		psp, _ := FromDomain_Psp(item)
		response.Results = append(response.Results, psp)
	}

	return &response, nil
}

type ListPsp struct {
	Results []Psp `json:"results" binding:"required"`
}

type PaginatedPspList struct {
	Results    []Psp                  `json:"results" binding:"required"`
	Pagination queryfilter.Pagination `json:"pagination" binding:"required"`
}

// ResponseFromDomain_Psp
func FromDomain_Psp(psp domain.PSP) (Psp, error) {
	fmt.Println(psp)
	pspResponse := Psp{
		ID:             *psp.ID, // If it will be nil - you're cannot dereference it (nil pointer dereference)
		PspCountryCode: psp.PspCountryCode,
		PspShortName:   psp.PspShortName,
		PspCode:        psp.PspCode,
		A:              FromDomain_AuditDBRecord(psp.AuditDBRecord),
	}

	return pspResponse, nil
}

// ResponseFromDomain_ListPsp
func FromDomainList_PspList(domainItems []domain.PSP) (*ListPsp, error) {

	//psps := []Psp{}

	listPsp := ListPsp{
		Results: []Psp{}, // Initialize with an empty slice
	}

	for _, psp := range domainItems {
		response, _ := FromDomain_Psp(psp)
		listPsp.Results = append(listPsp.Results, response)
	}

	return &listPsp, nil

}

func FromDomain_PaginatedPspList(domainItems []domain.PSP, p queryfilter.Pagination) (*PaginatedPspList, error) {

	//psps := []Psp{}

	response := PaginatedPspList{
		Results:    []Psp{}, // Initialize with an empty slice
		Pagination: p,
	}

	for _, item := range domainItems {
		psp, _ := FromDomain_Psp(item)
		response.Results = append(response.Results, psp)
	}

	return &response, nil

}
