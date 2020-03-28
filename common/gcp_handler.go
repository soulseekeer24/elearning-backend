// GCPHandler handler for google cloud function on google cloud
func GCPHandler(w http.ResponseWriter, r *http.Request) {
	// BodyRequest will be used to take the json response from client and build it
	bodyRequest := common.BodyRequest{
		Keywords: []string{},
	}

	err := json.NewDecoder(r.Body).Decode(&bodyRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	courseList, err := searchForCourse(bodyRequest.Keywords)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bodyResponse := common.BodyResponse{
		Courses: courseList,
	}

	response, err := json.Marshal(&bodyResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}
