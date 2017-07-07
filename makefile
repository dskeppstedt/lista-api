dir = $(PWD)

run:
	docker run --rm -it -p 8080:8080 -v $(dir)/src:/go/src/lista/api/ --name lista-api la-image go run main.go
build:
	docker build --no-cache -t la-image .