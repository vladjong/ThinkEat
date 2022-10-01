APP_BIN = build/app

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) cmd/main/main.go
	./build/app

clean:
	rm -rf build || true

swagger:
	swag init -g cmd/main/main.go -o ./docs

lint:
	golangci-lint run
