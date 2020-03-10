package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/chromedp"
	"github.com/golearn/common"
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

func HandleRequest(ctx context.Context, request common.BodyRequest) ([]common.CourseInfo, error) {
	// create context
	result := SearchEDX("lol")
	return result, nil
}

func SearchEDX(courseName string) []common.CourseInfo {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []common.CourseInfo
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
func main() {
	lambda.Start(HandleRequest)
}
