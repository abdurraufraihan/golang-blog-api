package utils

import "strings"

type Response struct {
	Status bool        `json:"status"`
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

func GetResponse(data interface{}) Response {
	res := Response{
		Status: true,
		Errors: nil,
		Data:   data,
	}
	return res
}

func GetErrorResponse(err string) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status: false,
		Errors: splittedError,
		Data:   struct{}{},
	}
	return res
}
