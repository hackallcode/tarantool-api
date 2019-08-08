package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"kv-storage/internal/pkg/logger"
	"kv-storage/internal/pkg/models"
)

func getStrUrlParam(w http.ResponseWriter, r *http.Request, name string) (string, bool) {
	vars := mux.Vars(r)
	result, ok := vars[name]
	if ok {
		return result, true
	} else {
		sendJson(w, http.StatusBadRequest, models.IncorrectParamAnswer)
		return "", false
	}
}

func parseJson(w http.ResponseWriter, r *http.Request, result models.InputModel) bool {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return false
	}

	if err = r.Body.Close(); err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return false
	}

	if err = result.UnmarshalJSON(data); err != nil {
		sendJson(w, http.StatusBadRequest, models.IncorrectJsonAnswer)
		return false
	}

	if !result.Validate() {
		sendJson(w, http.StatusBadRequest, models.IncorrectJsonAnswer)
		return false
	}

	return true
}

func handleCommonErrors(w http.ResponseWriter, err error) bool {
	if err != nil {
		if err == models.AlreadyExistsError {
			sendJson(w, http.StatusConflict, models.KeyExistsAnswer)
			return false
		}
		if err == models.NotFoundError {
			sendJson(w, http.StatusNotFound, models.KeyNotFoundAnswer)
			return false
		}
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return false
	}
	return true
}

func sendJson(w http.ResponseWriter, statusCode int, outModel models.OutputModel) {
	answer, err := outModel.MarshalJSON()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	w.WriteHeader(statusCode)
	_, err = fmt.Fprintln(w, string(answer))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
