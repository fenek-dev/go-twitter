package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, code int, message string, data interface{}) {

	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	resJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(code)
	w.Write(resJson)
}
