CLI_NAME = ism2

all: cli

cli:
	go build -o ${CLI_NAME} cmd/cli/main.go

