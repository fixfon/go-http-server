package controller

import (
	"encoding/json"
	"fmt"
	"http-server-go/service"
	"net/http"
)

func RandomElement(responseWriter http.ResponseWriter, request *http.Request) {
	element := service.RandomElement()

	err := json.NewEncoder(responseWriter).Encode(element)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func IndexElement(responseWriter http.ResponseWriter, request *http.Request) {
	var id int
	_, err := fmt.Sscanf(request.URL.Path, "/element/%d", &id)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	element, err := service.GetElementById(id)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(element)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
