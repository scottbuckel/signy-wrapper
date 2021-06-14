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
}

func SignyHandler(w http.ResponseWriter, r *http.Request) {

	var SignyReturn SignyReturn

	SignyReturn.FailureReason = ""
	SignyReturn.SignyValidation = "success"

	min := 1
	max := 30
	SignyReturn.RandomNumber = rand.Intn(max-min) + min

	if SignyReturn.RandomNumber%2 == 0 {
		SignyReturn.FailureReason = "Number was Even, Evens are failures"
		SignyReturn.SignyValidation = "failure"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(SignyReturn)
}
