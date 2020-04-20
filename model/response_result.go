package model

type ResponseResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
