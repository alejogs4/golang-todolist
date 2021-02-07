package infraestructure

import (
	"encoding/json"
	"net/http"
)

func DispatchNewHttpError(w http.ResponseWriter, message string, statusCode int) {
	responseContent, _ := json.Marshal(map[string]string{"message": message})

	w.WriteHeader(statusCode)
	w.Write(responseContent)
}

func WrapAPIResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
	}
}
