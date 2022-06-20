APP=backend_go
APP_EXE="./build/$(APP)"

run:
	$(APP_EXE) server
build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux go build -o ${APP_EXE}