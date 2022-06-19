APP=angga-backend-go
APP_EXE="./build/$(APP)"

run:
	$(APP_EXE) server
build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=windows go build -o ${APP_EXE}