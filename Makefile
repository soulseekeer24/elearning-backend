.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./platzi-scrapper/platzi-scrapper
	rm -rf ./edx-scrapper/edx-scrapper
	
build:
	GOOS=linux GOARCH=amd64 go build -o platzi-scrapper/platzi-scrapper ./platzi-scrapper
	GOOS=linux GOARCH=amd64 go build -o edx-scrapper/edx-scrapper ./edx-scrapper