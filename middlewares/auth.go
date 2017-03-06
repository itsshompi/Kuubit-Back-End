package middlewares

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	secretKey = "kuubitSecretKey"
)

//Message Struct
type Message struct {
	Message string `json:"message"`
}

//AuthMiddleware is ...
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err == nil && token.Valid {
		next(w, r)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			ResponseJSON("Token mal formateado", http.StatusForbidden, w)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			ResponseJSON("Token Expirado", http.StatusForbidden, w)
			return
		} else {
			ResponseJSON("Error", http.StatusForbidden, w)
			return
		}
	}
}

//ResponseJSON is ...
func ResponseJSON(message string, status int, w http.ResponseWriter) {
	json, e := json.Marshal(Message{message})
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	return
}
