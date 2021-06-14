package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type SignyReturn struct {
	SignyValidation string `json:"SignyValidation"`
	FailureReason   string `json:"FailureReason"`
	RandomNumber    int    `json:"RandomNumber"`
	ImageName       string `json:"ImageName"`
}

func SignyHandler(w http.ResponseWriter, r *http.Request) {

	var SignyReturn SignyReturn

	SignyReturn.FailureReason = ""
	SignyReturn.SignyValidation = "failure"

	keys, ok := r.URL.Query()["image"]

	if !ok || len(keys[0]) < 1 {
		SignyReturn.FailureReason = "No Image Supplied"
		SignyReturn.SignyValidation = "failure"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(SignyReturn)
	} else {
		SignyReturn.ImageName = keys[0]
		min := 1
		max := 100
		SignyReturn.RandomNumber = rand.Intn(max-min) + min

		if SignyReturn.RandomNumber%2 == 0 {
			SignyReturn.FailureReason = "Number was Even, Evens are failures"
			SignyReturn.SignyValidation = "failure"
		} else {
			SignyReturn.FailureReason = ""
			SignyReturn.SignyValidation = "success"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(SignyReturn)
	}

}
