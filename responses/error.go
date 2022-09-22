package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// var global
var (
	OkStatus PayloadStatus = PayloadStatus{Code: "SUCCESS", Message: ""}
)

// PayloadStatus is class
type PayloadStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ParseStatus is func init status
func ParseStatus(code string, message string) PayloadStatus {
	return PayloadStatus{
		Code:    code,
		Message: message,
	}
}

// HasError is func test err
func (e PayloadStatus) HasError() bool {
	if e.Code == "SUCCESS" {
		return false
	}
	return e.Code != ""
}
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
