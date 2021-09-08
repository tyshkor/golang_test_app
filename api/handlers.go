package api

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/tyshkor/golang_test_app/model"
	"github.com/tyshkor/golang_test_app/usecase"
)

func HandlerAPI1(w http.ResponseWriter, r *http.Request) {

	payload := model.API1Model{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := model.ConvertAPI1ToStandard(payload)

	res := usecase.PayloadHandler(data)

	output := model.OutPutModelAPI1{
		Total: res,
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerAPI2JSON(w http.ResponseWriter, r *http.Request) {

	payload := model.API2Model{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := model.ConvertAPI2ToStandard(payload)

	res := usecase.PayloadHandler(data)

	output := model.OutPutModelAPI2{
		Amount: res,
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerAPI2XML(w http.ResponseWriter, r *http.Request) {

	payload := model.API2Model{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := model.ConvertAPI2ToStandard(payload)

	res := usecase.PayloadHandler(data)

	output := model.OutPutModelAPI2{
		Amount: res,
	}

	if err := xml.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
