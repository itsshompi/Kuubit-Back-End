package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type infoApp struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository string `json:"repository"`
	CreatedBy  string `json:"created_by"`
	Github     string `json:"github"`
}

//HomeController is ...
func HomeController(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HomeController")
	w.Header().Set("Content-Type", "application/json")
	info := infoApp{
		"Kuubit API",
		"1.0.1",
		"https://github.com/itsshompi/kuubit-backend",
		"Felipe Schneeberger",
		"https://github.com/itsshompi",
	}
	if json, err := json.MarshalIndent(info, "", "    "); err == nil {
		w.Write(json)
	}
}
