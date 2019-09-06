package utils

import (
	"encoding/json"
	"net/http"
)

// type Result struct {
// 	Message string      `json:"message"`
// 	Status  bool        `json:"status"`
// 	Data    interface{} `json:"data"`
// }

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    status,
		"message": message,
	}
}

func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(data)
}

// func Jawab(status bool, message string, data map[string]interface{}) *Result {
// 	ret := Result{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}
// 	return &ret
// }

// func Siap(w http.ResponseWriter, data *Result) {
// 	w.Header().Add("Content-Type", "aplication/json")
// 	json.NewEncoder(w).Encode(data)
// }
