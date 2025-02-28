// domainerror has a custom error type
package domainerror

import "errors"

type ValidationError struct {
	Err              error
	ValidationErrors []error
}

func (r *ValidationError) Error() string {
	return r.Err.Error()
}

var (
	ErrValidationError                 = errors.New("validation error")
	ErrInvalidInputData                = errors.New("invalid input data")
	ErrNameTooShort                    = errors.New("[name] is too short")
	ErrNameTooLong                     = errors.New("[name] is too long")
	ErrNotValidIsoCode                 = errors.New("[isoCode] not valid, the expected length is 2")
	ErrBadFilterField                  = errors.New("unsupported filter field")
	ErrSortFieldNotInPair              = errors.New("sorting field needs to be in form ['country','ASC']")
	ErrBadSortField                    = errors.New("unsupported sort field")
	ErrBadSortDirectionField           = errors.New("sorting direction can be either 'ASC' or 'DESC'")
	ErrRangeListNotAPair               = errors.New("range list needs to have 2 elements, eg: [1,42]")
	ErrRangeListSecondElementNotBigger = errors.New("range list second element needs to be bigger or equal to the first element, eg: [1,42] - 42 is bigger than 1")
	ErrDuplicateDataCreate             = errors.New("the data you are trying to create already exists in the database")
	ErrNoEntryFoundInDb                = errors.New("no entry found in db")
)
