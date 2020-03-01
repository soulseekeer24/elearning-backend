.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./platzi-scrapper/platzi-scrapper
	
build:
	GOOS=linux GOARCH=amd64 go build -o platzi-scrapper/platzi-scrapper ./platzi-scrapper