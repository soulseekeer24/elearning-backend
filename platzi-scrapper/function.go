package platzi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golearn/common"
)

const baseURL = "https://platzi.com"
const courseListSelector = "body > section.SearcherMaterial > div > ul > a"
const courseImgSelector = ".SearcherMaterial-itemImage > img"
const courseTitleSelector = ".SearcherMaterial-itemName"

func searchForCourse(courses []string) ([]common.CourseInfo, error) {

	var courseQuery = ""
	for i := 0; i < len(courses); i++ {
		courseQuery += courses[i] + "+"
	}

	resp, err := http.Get(baseURL + "/search/?search=" + courseQuery + "&filter=course")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	coursesList := make([]common.CourseInfo, 0)

	doc.Find(courseListSelector).Each(func(i int, s *goquery.Selection) {

		imageURL, _ := s.Find(courseImgSelector).First().Attr("src")
		link, _ := s.Attr("href")

		course := common.CourseInfo{
			Title:    strings.TrimSpace(s.Find(courseTitleSelector).First().Text()),
			ImageURL: strings.TrimSpace(imageURL),
			URL:      strings.TrimSpace(baseURL + link)}

		coursesList = append(coursesList, course)
	})

	return coursesList, nil
}

// HandlerPlatzi handler for google cloud function on google cloud
func HandlerPlatzi(w http.ResponseWriter, r *http.Request) {
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
