package edx

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

const API = "https://www.edx.org/es/course?search_query="

const jssQuery = `
		(function(){ 
			let cards=document.querySelectorAll(".discovery-card");
			let data = [];

			cards.forEach(c=>{
				let link = c.querySelector("a").getAttribute("href");
				let titleEle = c.querySelector("div.title.ellipsis-multi-line h3.title-heading.ellipsis-overflowing-child");    
				let title;
				let imageEle = c.querySelector("a.course-link div.img-wrapper img");
				let d= {};
				d['URL']=link;
				if(imageEle){
					d['imageUrl'] = imageEle.src;
				}
				
				if(titleEle){
					d['title'] = titleEle.textContent;
				}
				data.push(d);
			});
			return data;
		})();
`

type CourseInfo struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	ImageURL string `json:"imageURL"`
}

func HandleRequest(ctx context.Context, courseName string) ([]CourseInfo, error) {
	// create context
	result := SearchEDX(courseName)
	return result, nil
}

func SearchEDX(courseName string) []CourseInfo {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []CourseInfo
	err := chromedp.Run(ctx,
		chromedp.Navigate(API+courseName),
		chromedp.WaitVisible(`#search-results-section`, chromedp.ByID),
		chromedp.Evaluate(jssQuery, &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
