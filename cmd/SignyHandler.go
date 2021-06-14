package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/scottbuckel/signy-wrapper/version"
)

func SignyHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var requestGun RequestGun
	json.Unmarshal(reqBody, &requestGun)

	var info Info

	info.NotaryWrapperVersion = version.NotaryWrapperVersion
	info.GitCommit = version.GitCommit
	info.RuntimeVersion = runtime.Version()

	info.SignyValidation = "success"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(info)
}
