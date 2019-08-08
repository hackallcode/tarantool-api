package controllers

import (
	"net/http"

	"kv-storage/internal/pkg/models"
	"kv-storage/internal/pkg/pair"
)

func CreatePair(w http.ResponseWriter, r *http.Request) {
	input := models.CreatePairData{}
	if !parseJson(w, r, &input) {
		return
	}

	err := pair.CreatePair(input)
	if !handleCommonErrors(w, err) {
		return
	}

	sendJson(w, http.StatusOK, models.PairCreatedAnswer)
}

func GetPair(w http.ResponseWriter, r *http.Request) {
	key, exists := getStrUrlParam(w, r, "key")
	if !exists {
		return
	}

	output, err := pair.GetPair(key)
	if !handleCommonErrors(w, err) {
		return
	}

	sendJson(w, http.StatusOK, output)
}

func UpdatePair(w http.ResponseWriter, r *http.Request) {
	key, exists := getStrUrlParam(w, r, "key")
	if !exists {
		return
	}

	input := models.UpdatePairData{}
	if !parseJson(w, r, &input) {
		return
	}

	err := pair.UpdatePair(key, input)
	if !handleCommonErrors(w, err) {
		return
	}

	sendJson(w, http.StatusOK, models.PairUpdatedAnswer)
}

func RemovePair(w http.ResponseWriter, r *http.Request) {
	key, exists := getStrUrlParam(w, r, "key")
	if !exists {
		return
	}

	err := pair.RemovePair(key)
	if !handleCommonErrors(w, err) {
		return
	}

	sendJson(w, http.StatusOK, models.PairRemovedAnswer)
}
