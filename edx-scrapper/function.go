package edxscrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golearn/common"
)

const baseURL = "https://edx.org/api/v1/catalog"

type EDXQueryResponse struct {
	Objects EDXObject
}

type EDXObject struct {
	Results []EDXListingCourse
}

type EDXListingCourse struct {
	FullDescription string `json:"full_description"`
	LevelType       string `json:"level_type "`
	Title           string `json:"title"`
	ImageURL        string `json:"image_url"`
	MarketingURL    string `json:"marketing_url"`
}

func searchForCourse(courses []string) ([]common.CourseInfo, error) {

	var courseQuery = ""
	coursesList := make([]common.CourseInfo, 0)
	for i := 0; i < len(courses); i++ {
		courseQuery += courses[i]
	}
	fullUrl := baseURL + "/search/?query=" + courseQuery
	fmt.Println(fullUrl)

	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data EDXQueryResponse
	json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	for _, courseEdx := range data.Objects.Results {
		course := common.CourseInfo{
			Title:    strings.TrimSpace(courseEdx.Title),
			ImageURL: strings.TrimSpace(courseEdx.ImageURL),
			URL:      strings.TrimSpace(courseEdx.MarketingURL)}

		coursesList = append(coursesList, course)
	}

	return coursesList, nil
}

// HandlerEDX handler for google cloud function on google cloud
func HandlerEDX(w http.ResponseWriter, r *http.Request) {
	common.GCPHandler(w, r, searchForCourse)
}
