package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/scottbuckel/signy-wrapper/version"
)

type SignyReturn struct {
	NotaryWrapperVersion string `json:"NotaryWrapperVersion"`
	GitCommit            string `json:"GitCommit"`
	RuntimeVersion       string `json:"runtimeVersion"`
	SignyValidation      string `json:"SignyValidation"`
}

func SignyHandler(w http.ResponseWriter, r *http.Request) {

	var SignyReturn SignyReturn

	SignyReturn.NotaryWrapperVersion = version.NotaryWrapperVersion
	SignyReturn.GitCommit = version.GitCommit
	SignyReturn.RuntimeVersion = runtime.Version()

	SignyReturn.SignyValidation = "failure"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(SignyReturn)
}
