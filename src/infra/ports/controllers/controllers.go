package controllers

import (
	shared_app "abstract-entities/src/shared/app"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type StatusReturn struct {
	Status bool `json:"status"`
}

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	HealtCheck := StatusReturn{
		Status: true,
	}
	data, err := json.Marshal(HealtCheck)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func CreateCNH(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	var body shared_app.CNH
	err = json.Unmarshal(data, &body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	service, err := shared_app.GetServiceCNH(body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	output, errors := service.Create()
	if len(errors) > 0 {
		output, _ := json.Marshal(errors)
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Set("content-type", "application/json")
		res.Write(output)
		return
	}
	outputResponse, _ := json.Marshal(output)
	res.WriteHeader(http.StatusOK)
	res.Header().Set("content-type", "application/json")
	res.Write(outputResponse)
	return
}
