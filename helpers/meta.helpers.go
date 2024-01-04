package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"skeleton-svc/constants"
	"skeleton-svc/helpers/models"
	"strings"
)

// ListErrorCode ...
var (
	ListErrorCode []models.MappingErrorCodes
)

// RegisterErrorCode registering code
func RegisterErrorCode() bool {

	var b []byte
	var err error
	b, err = os.ReadFile("errorcodes.json") // just pass the file name
	if err != nil {
		b, err = os.ReadFile("../errorcodes.json") // just pass the file name
		if err != nil {
			fmt.Println("Failed to read file error code json : ", err)
		}
	}

	if json.Unmarshal(b, &ListErrorCode) != nil {
		fmt.Println("Failed to unmarshaling json response error mapping ")
		return false
	}

	return true
}

// GetMetaResponse ..
func GetMetaResponse(key string) models.MetaData {

	// fmt.Println("Get meta response by key:", key)

	var meta models.MetaData

	if key == constants.RcSuccess {
		meta.Code = fmt.Sprintf("200%v00", constants.SERVICE_CODE)
		meta.Title = "Success"
		meta.Message = "Successful"
		return meta
	}

	for _, element := range ListErrorCode {
		if element.Key == key {
			meta.Code = strings.Replace(element.Content.Code, "?", constants.SERVICE_CODE, 1)
			meta.Title = element.Content.Title
			meta.Message = element.Content.Message
			return meta
		}
	}

	meta.Code = fmt.Sprintf("%v%v%v", http.StatusInternalServerError, constants.SERVICE_CODE, "00")
	meta.Title = "Error"
	meta.Message = "General error"
	return meta
}
