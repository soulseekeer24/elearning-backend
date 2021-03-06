package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const baseURL = "https://platzi.com"
const courseListSelector = "body > section.SearcherMaterial > div > ul > a"
const courseImgSelector = ".SearcherMaterial-itemImage > img"
const courseTitleSelector = ".SearcherMaterial-itemName"

// CourseInfo represent all the data for course
type CourseInfo struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	ImageURL string `json:"imageURL"`
}

// BodyRequest represent all the data for course
type BodyRequest struct {
	Keywords []string `json:"keywords"`
}

// BodyResponse represent all the data for course
type BodyResponse struct {
	Courses []CourseInfo `json:"courses"`
}

func searchForCourse(courses []string) ([]CourseInfo, error) {

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

	coursesList := make([]CourseInfo, 0)

	doc.Find(courseListSelector).Each(func(i int, s *goquery.Selection) {

		imageURL, _ := s.Find(courseImgSelector).First().Attr("src")
		link, _ := s.Attr("href")

		course := CourseInfo{
			Title:    strings.TrimSpace(s.Find(courseTitleSelector).First().Text()),
			ImageURL: strings.TrimSpace(imageURL),
			URL:      strings.TrimSpace(baseURL + link)}

		coursesList = append(coursesList, course)
	})

	return coursesList, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// BodyRequest will be used to take the json response from client and build it
	bodyRequest := BodyRequest{
		Keywords: []string{},
	}

	err := json.Unmarshal([]byte(request.Body), &bodyRequest)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	courseList, err := searchForCourse(bodyRequest.Keywords)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	bodyResponse := BodyResponse{
		Courses: courseList,
	}

	response, err := json.Marshal(&bodyResponse)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	resp := events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}

	resp.Headers = make(map[string]string)
	resp.Headers["Content-Type"] = "application/json"
	return resp, nil
}

func main() {
	lambda.Start(handler)
}
