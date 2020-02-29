package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://platzi.com"

// CourseInfo represent all the data for course
type CourseInfo struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	ImageURL string `json:"imageURL"`
}

func main() {

	coursesScrapped, err := searchForCourse("web")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("course = %v \n", coursesScrapped)
}

// Search in platzi.com for any course with the given name
func searchForCourse(courses string) ([]CourseInfo, error) {
	resp, err := http.Get(baseURL + "/search/?search=" + courses + "&filter=course")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	coursesList := make([]CourseInfo, 1)

	doc.Find("body > section.SearcherMaterial > div > ul > a").Each(func(i int, s *goquery.Selection) {

		imageURL, existImage := s.Find(".SearcherMaterial-itemImage > img").First().Attr("src")
		link, existLink := s.Attr("href")

		if existImage && existLink {
			course := CourseInfo{
				Title:    s.Find(".SearcherMaterial-itemName").First().Text(),
				ImageURL: imageURL,
				URL:      baseURL + link}

			coursesList = append(coursesList, course)
		}
	})

	return coursesList, nil
}
