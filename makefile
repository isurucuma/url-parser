BINARY_NAME=myapp

.PHONY: build run clean test

build:
	@echo "Building..."
	go build -o ${BINARY_NAME} main.go

run: build
	@echo "Running..."
	./${BINARY_NAME}

clean:
	@echo "Cleaning..."
	go clean
	rm -f ${BINARY_NAME}

test:
	@echo "Testing..."
	go test -v ./...

