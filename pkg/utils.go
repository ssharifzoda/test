package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, text string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(text))
}

func Response(w http.ResponseWriter, data interface{}) {
	date, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(date)
}
