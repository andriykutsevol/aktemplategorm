package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// ItemsToContentRangeString - converts a series of values into a content-range string, to be returned in response header
func ItemsToContentRangeString(entityName string, safeSubsliceIndex0, safeSubsliceIndex1, totalItems uint) string {
	//due to the fact that we previously increased the secondIndex by 1, in LimitSubSliceIndexes, we need to decrease it when sending
	//the range information to FE, te be in sync with the data they sent us in "range" query parameter
	return fmt.Sprintf("%s %v-%v/%v", entityName, safeSubsliceIndex0, safeSubsliceIndex1-1, totalItems)
}

//	LimitSubSliceIndexes - limits subslice indexes so that when subslicing the target slice, we don't get an out of bounds error
//
// safeSubsliceIndex0, safeSubsliceIndex1 can safely be used to subslice the targetSlice, like targetSlice[safeSubsliceIndex0:safeSubsliceIndex1]
func LimitSubSliceIndexes(subSliceIndex0, subSliceIndex1 uint, targetSliceLen uint) (safeSubsliceIndex0, safeSubsliceIndex1 uint) {
	safeSubsliceIndex0 = subSliceIndex0
	//Due do the fact that FE expects inclusive last item (eg: [0,2], where 2 is inclusive) AND
	//due to the fact that in golang, when sublicing, the last element is exclusive (eg: s[0,2], where 2 is exclusive )
	//we need to increase the last item in the slice range by 1
	safeSubsliceIndex1 = subSliceIndex1 + 1

	if safeSubsliceIndex1 > targetSliceLen {
		safeSubsliceIndex1 = targetSliceLen
	}
	if safeSubsliceIndex0 > safeSubsliceIndex1 {
		safeSubsliceIndex0 = safeSubsliceIndex1
	}
	return safeSubsliceIndex0, safeSubsliceIndex1
}

// jsonBodyToReqMap - creates a request map (a map[string]interface) which is used to selectively update/create fields in an entity corporateSubsidiary
func JsonBodyToReqMap(r http.Request) (map[string]interface{}, error) {

	var reqMap map[string]interface{}
	jsonData, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &reqMap)

	if err != nil {
		return nil, err
	}

	return reqMap, nil
}

func IsValidJSON(jsonData []byte) bool {
	var decoded interface{}
	err := json.Unmarshal(jsonData, &decoded)
	return err == nil
}

func IsValidBodyJson(c *gin.Context) bool {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return IsValidJSON(bodyBytes)
}

// TODO: Move this to helpers
func StringToUint(str string) (uint, error) {
	// Convert the string to a uint64 to handle potential overflow

	if strings.HasPrefix(str, "-") {
		return 0, errors.New("negative numbers are not allowed for unsigned integers")
	}

	uint64Value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err // Return an error if the conversion fails
	}
	// Check if the uint64 value fits within the uint range
	if uint64Value > math.MaxUint32 {
		return 0, errors.New("value overflows uint") // Return an error if overflow occurs
	}
	return uint(uint64Value), nil // Return the converted uint value and no error
}
