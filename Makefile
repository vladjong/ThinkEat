APP_BIN = app/build/app

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

clean:
	rm -rf ./app/build || true

swagger:
	swag init -g ./app/cmd/app/main.go -o ./app/docs

lint:
	golangci-lint run
