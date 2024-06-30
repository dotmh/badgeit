BINARY_NAME:="badgeit"
BINARY_PATH:="dist"

build:
    go build -o {{BINARY_PATH}}/{{BINARY_NAME}}

run:
    go run main.go

tidy:
    go mod tidy