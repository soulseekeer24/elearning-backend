package udemy

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golearn/common"
)

const baseURL = "https://www.udemy.com/api-2.0/"
const clientID = "pBZRtd8x6r4rcVQL5QyZwoQvwsZTZYZMLGMn4y9E"
const clientSecret = "PekiJF6uIDpNabKKWPVN5EsqcXPPAvp43NuY8ry4AED4BFkaJiVyXWVSiubwqWkuE7zpicBS2RFd8fpYnxxfWEEJbSWSOgyd5XH0W2qxfjkfL92z16MFclx9AJD41qL0"

// ResponseUdemy response object from udemy api 2.0
type ResponseUdemy struct {
	Count    int32         `json:"count"`
	Next     string        `json:"next "`
	Previous string        `json:"previous"`
	Results  []UdemyCourse `json:"results"`
}

// UdemyCourse udemy course info
type UdemyCourse struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Price    string `json:"price"`
	Image    string `json:"image_240x135"`
	Headline string `json:"headline"`
}

// HandlerUdemy handler for google cloud function on google cloud
func HandlerUdemy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

func searchForCourse(courses []string) ([]common.CourseInfo, error) {

	var courseQuery = ""
	for i := 0; i < len(courses); i++ {
		courseQuery += courses[i] + " "
	}
	request, _ := http.NewRequest("GET", baseURL+"courses/?search="+courseQuery, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+b64.URLEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data ResponseUdemy
	json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	coursesList := make([]common.CourseInfo, 0)

	for _, courseUdemy := range data.Results {
		course := common.CourseInfo{
			Title:    strings.TrimSpace(courseUdemy.Title),
			ImageURL: strings.TrimSpace(courseUdemy.Image),
			URL:      strings.TrimSpace("https://www.udemy.com" + courseUdemy.URL)}

		coursesList = append(coursesList, course)
	}
	return coursesList, nil
}
