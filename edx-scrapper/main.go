package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
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

func HandleRequest(ctx context.Context, request common.BodyRequest) (events.APIGatewayProxyResponse, error) {
	// create context
	result := SearchEDX("lol")

	bodyResponse := common.BodyResponse{
		Courses: result,
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
