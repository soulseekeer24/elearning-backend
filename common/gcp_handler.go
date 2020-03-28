package common

import (
	"encoding/json"
	"net/http"
)

// GCPHandler handler for google cloud function on google cloud
func GCPHandler(w http.ResponseWriter, r *http.Request, scrapper ScrapperFunction) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// BodyRequest will be used to take the json response from client and build it
	bodyRequest := BodyRequest{
		Keywords: []string{},
	}

	err := json.NewDecoder(r.Body).Decode(&bodyRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	courseList, err := scrapper(bodyRequest.Keywords)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bodyResponse := BodyResponse{
		Courses: courseList,
	}

	response, err := json.Marshal(&bodyResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
