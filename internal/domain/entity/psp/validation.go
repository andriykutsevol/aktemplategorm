package psp

import (
	"errors"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/domain_error"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/patch"
	"github.com/andriykusevol/aktemplategorm/internal/domain/valueobject/queryfilter"
	"slices"
)

func ValidateFilterFields(filterFields queryfilter.FilterFields) error {
	approvedMapKeys := []string{"IDs", "PspCode", "PspCountryCode"}
	existingMapKeys := []string{}

	for _, v := range filterFields {
		existingMapKeys = append(existingMapKeys, v.Key)
	}

	//var e []error
	var e error
	//if we have a map key which is not present in the approved map keys return an err
	for _, v := range existingMapKeys {
		if !slices.Contains(approvedMapKeys, v) {
			if e != nil {
				e = fmt.Errorf("%v; %v; %v", e, domainerror.ErrBadFilterField, errors.New(v))
			} else {
				e = fmt.Errorf("%v; %v", domainerror.ErrBadFilterField, errors.New(v))
			}
		}
	}
	// if len(e) > 0 {
	// 	return errors.Join(e...)
	// }
	if e != nil {
		return e
	}
	return nil
}

func ValidateOrderFields(orderFields queryfilter.OrderFields) error {

	approvedMapKeys := []string{"ID", "PspCode", "PspCountryCode"}
	existingMapKeys := []string{}

	for _, v := range orderFields {
		existingMapKeys = append(existingMapKeys, v.Key)
	}

	//var e []error
	var e error
	//if we have a map key which is not present in the approved map keys return an err
	for _, v := range existingMapKeys {
		if !slices.Contains(approvedMapKeys, v) {
			if e != nil {
				e = fmt.Errorf("%v; %v; %v", e, domainerror.ErrBadSortField, errors.New(v))
			} else {
				e = fmt.Errorf("%v; %v", domainerror.ErrBadSortField, errors.New(v))
			}
		}
	}
	// if len(e) > 0 {
	// 	return errors.Join(e...)
	// }
	if e != nil {
		return e
	}
	return nil
}

func ValidatePatch(data map[string]interface{}) bool {

	// for key, value := range data {
	// 	//TODO: Manually check for required fields, and business logic.
	// 	_ = key
	// 	_ = value
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	return true
}

func ToDomain(p *patch.Patch) error {
	return nil
}
