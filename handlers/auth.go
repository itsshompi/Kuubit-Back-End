package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "kuubitSecretKey"
)

type userCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//UserResponse Struct
type UserResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
}

//Message Struct
type Message struct {
	Message string `json:"message"`
}

//LoginHandler ....
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user userCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println(user.Username, user.Password)
	if strings.ToLower(user.Username) == "shompi" && user.Password == "1234567" {
		claims := jwt.MapClaims{
			"id":      "12AG2JDS92K9L923_1JHG1",
			"name":    "Felipe Schneeberger",
			"picture": "https://buckets3.kuubit.com/pictures/avatar/picture.jpg",
			"exp":     time.Now().Add(time.Minute * 20).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("Error signing token: %v\n", err)
			return
		}
		json, err := json.Marshal(UserResponse{t, user.Username, user.Password})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	}
	w.WriteHeader(http.StatusForbidden)
	fmt.Println("Invalid credentials")
	fmt.Fprint(w, "Invalid credentials")
	return
}

//AuthHandler is ....
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	fmt.Println(tokenString)
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		json, e := json.Marshal(Message{"Token válido"})
		if e != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("%v %v", claims["name"], claims["picture"])
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			json, e := json.Marshal(Message{"Token mal formateado"})
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write(json)
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			json, e := json.Marshal(Message{"Token expirado"})
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write(json)
			return
		} else {
			json, e := json.Marshal(Message{"Error"})
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write(json)
			return
		}
	} else {
		json, e := json.Marshal(Message{"Error"})
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(json)
		return
	}
}
