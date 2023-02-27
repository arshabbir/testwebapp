package api

import (
	"encoding/json"
	"log"
	"net/http"
	"testwebservermod/utils"
)

func (s *server) handlePing(w http.ResponseWriter, r *http.Request) {
	resp := utils.ApiError{Code: http.StatusOK, Message: "pong"}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Fatal("error while encoding.")
	}
}
