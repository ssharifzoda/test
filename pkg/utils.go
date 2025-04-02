package pkg

import (
	"encoding/json"
	"github.com/spf13/viper"
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

func InitConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
