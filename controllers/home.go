package controllers

import (
	"encoding/json"
	"net/http"
)

type infoApp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Website     string `json:"website"`
	URL         string `json:"url"`
	Repository  string `json:"repository"`
	CreatedBy   string `json:"created_by"`
	Github      string `json:"github"`
	Contact     string `json:"contact"`
}

//HomeController is ...
func HomeController(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := infoApp{
		"Kuubit API",
		"This application addresses to connect different applications of Kuubit (Web, Mobile, Desktop)",
		"www.kuubit.com",
		"api.kuubit.com",
		"1.0.1",
		"https://github.com/itsshompi/kuubit-backend",
		"Felipe Schneeberger",
		"https://github.com/itsshompi",
		"contact@kuubit.com",
	}
	if json, err := json.MarshalIndent(info, "", "    "); err == nil {
		w.Write(json)
	}
}
