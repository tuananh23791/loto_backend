package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"travel/model"
)

func responseToUser(errorCode int, errorMessage string, data interface{}, w http.ResponseWriter) {
	var res model.ResponseResult
	res.ErrorCode = errorCode
	res.Message = errorMessage
	res.Data = data
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println("here 6  " + err.Error())
	}
}

func getTokenString(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return ""
	}

	reqToken = strings.TrimSpace(splitToken[1])
	return reqToken
}
