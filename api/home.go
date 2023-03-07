package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"testwebservermod/models"
	"testwebservermod/utils"
)

func (s *server) handleHome(w http.ResponseWriter, r *http.Request) {

	d := &models.TemplateData{}
	render(w, r, "index.gohtml", d, s.conf.TemplatePath)
}

func render(w http.ResponseWriter, r *http.Request, tName string, data *models.TemplateData, templPath string) error {
	parsedFile, err := template.ParseFiles(templPath + tName)
	if err != nil {
		apiError := utils.ApiError{Code: http.StatusInternalServerError, Message: "unable to parse the files"}
		json.NewEncoder(w).Encode(*&apiError)
	}
	data.IP = getIPFromContext(r.Context())
	err = parsedFile.Execute(w, data)
	if err != nil {
		apiError := utils.ApiError{Code: http.StatusInternalServerError, Message: "unable to parse the files"}
		json.NewEncoder(w).Encode(*&apiError)
	}

	return err
}
