package controllers

import (
	"net/http"

	"kv-storage/internal/pkg/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, models.GetSuccessAnswer("Backend of project 'kv-storage'!"))
}
