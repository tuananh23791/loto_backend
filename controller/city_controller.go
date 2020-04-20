package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	ErrorCode "travel/config"
	"travel/model"
	"travel/repository"
)

func CreateCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cityList []model.City
	var cityListString []string
	err := r.ParseForm()
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	// decoder := schema.NewDecoder()
	fmt.Println("city name", r.PostForm.Get("name"))
	cityListName := r.PostForm.Get("name")
	// err = decoder.Decode(&cityListString, r.PostForm)
	// decoder := json.NewDecoder(strings.NewReader(cityListName))
	// err = decoder.Decode(&cityListString)
	err = json.Unmarshal([]byte(cityListName), &cityListString)
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}
	fmt.Println("cityListString ", cityListString)

	for _, v := range cityListString {
		var city = model.City{
			Name:      v,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
		fmt.Println("city -- ", city.Name)
		cityList = append(cityList, city)
	}

	reqToken := getTokenString(r)
	fmt.Println("reqToken ", reqToken)

	errorCode, errorMessage := repository.InsertCity(cityList, reqToken)
	if errorCode != ErrorCode.SUCCESS {
		responseToUser(errorCode, errorMessage, "", w)
		return
	}

	responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, "", w)
}

func GetCityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cityList []model.City
	err := r.ParseForm()
	if err != nil {
		responseToUser(ErrorCode.SOME_THING_WENT_WRONG, err.Error(), "", w)
		return
	}

	reqToken := getTokenString(r)
	fmt.Println("reqToken ", reqToken)

	errorCode, errorMessage, cityList := repository.GetListCity(reqToken)
	if errorCode != ErrorCode.SUCCESS {
		responseToUser(errorCode, errorMessage, "", w)
		return
	}

	if len(cityList) == 0 {
		responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, "", w)
	} else {
		responseToUser(ErrorCode.SUCCESS, ErrorCode.SUCCESS_MESSAGE, cityList, w)
	}
}
