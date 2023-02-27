package api

import (
	"encoding/json"
	"net/http"
	"testwebservermod/models"
	"testwebservermod/utils"
	"text/template"
)

func (s *server) handleHome(w http.ResponseWriter, r *http.Request) {

	d := &models.TemplateData{}
	render(w, r, "index.gohtml", d)
}

func render(w http.ResponseWriter, r *http.Request, tName string, data *models.TemplateData) error {
	parsedFile, err := template.ParseFiles("./templates/" + tName)
	if err != nil {
		apiError := utils.ApiError{Code: http.StatusInternalServerError, Message: "unable to parse the files"}
		json.NewEncoder(w).Encode(*&apiError)
	}
	err = parsedFile.Execute(w, data)
	if err != nil {
		apiError := utils.ApiError{Code: http.StatusInternalServerError, Message: "unable to parse the files"}
		json.NewEncoder(w).Encode(*&apiError)
	}

	return err
}
