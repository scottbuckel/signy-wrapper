package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/scottbuckel/signy-wrapper/version"
)

type Info struct {
	NotaryWrapperVersion string `json:"NotaryWrapperVersion"`
	GitCommit            string `json:"GitCommit"`
	RuntimeVersion       string `json:"runtimeVersion"`
	SignyValidation      string `json:"SignyValidation"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var info Info

	info.NotaryWrapperVersion = version.NotaryWrapperVersion
	info.GitCommit = version.GitCommit
	info.RuntimeVersion = runtime.Version()

	info.SignyValidation = "success"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(info)
}
