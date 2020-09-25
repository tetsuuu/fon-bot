.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./notify/fondesk-bot
	
build:
	GOOS=linux GOARCH=amd64 go build -o notify/fondesk-bot ./notify
