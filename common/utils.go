package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	appMessage struct {
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
		Code       int    `json:"code"`
	}
	messageResource struct {
		Data appMessage `json:"data"`
	}
	configuration struct {
		Server      string
		MongoDBHost string
		DBUser      string
		DBPwd       string
		Database    string
		LogLevel    int
	}
)

//DisplayAppError ...
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, status int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: status,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

//DisplayAppMessage is...
func DisplayAppMessage(w http.ResponseWriter, message string, status int, code int) {
	msgObj := appMessage{
		Message:    message,
		HTTPStatus: status,
		Code:       code,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if j, err := json.Marshal(messageResource{Data: msgObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
