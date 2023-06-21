TARGET=gomatching.thienht
DIR=build
$(warning *******************************************************)
all: clean build
init:
	export PATH=$PATH:/usr/local/go/bin
doc:
	./swag init -g pkg/controller/*
run:
	go run cmd/dev/main.go
clean: 
	rm -rf $(DIR)
build_pro:
	go vet
	go fmt
	go build -o $(DIR)/$(TARGET) main.go
	scp $(DIR)/$(TARGET) :'/home/thienhang/gochat/core'

