package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	ErrorCode "travel/config"
	"travel/model"
	"travel/repository"

	"github.com/gorilla/schema"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := r.ParseForm()
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&user, r.PostForm)
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	errorCode, errorMessage, responseUser := repository.InsertUser(user)
	if errorCode != ErrorCode.SUCCESS {
		responseToUser(errorCode, errorMessage, "", w)
		return
	}

	responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, responseUser, w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := r.ParseForm()
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&user, r.PostForm)
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	errorCode, errorMessage, responseUser := repository.Login(user)
	if errorCode != ErrorCode.SUCCESS {
		responseToUser(errorCode, errorMessage, "", w)
		return
	}

	responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, responseUser, w)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := r.ParseForm()
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&user, r.PostForm)
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
	}

	reqToken = strings.TrimSpace(splitToken[1])

	errorCode, errorMessage, responseUser := repository.UpdateUser(user, reqToken)

	if errorCode != ErrorCode.SUCCESS {
		responseToUser(errorCode, errorMessage, "", w)
		return
	}

	responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, responseUser, w)
}

func convertStructToString(data interface{}) (int, string) {
	b, err := json.Marshal(data)
	fmt.Println("here 3  ")
	if err != nil {
		fmt.Println("here 4  ", err.Error())
		fmt.Println(err)
		return ErrorCode.SOME_THING_WENT_WRONG, err.Error()
	}

	return ErrorCode.SUCCESS, string(b)
}
