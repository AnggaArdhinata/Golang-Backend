APP=backend_go
APP_EXE="./build/$(APP)"

runs:
	go run main.go server

run:
	./build/$(APP) server
build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_EXE}