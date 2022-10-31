SHELL = /bin/bash

SRC_FOLDER=.
NAME = demail-proto
MAIN = $(SRC_FOLDER)/demail.proto
CC = protoc

.PHONY: all
all: compile

.PHONY: compile
compile:
	$(CC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(MAIN)