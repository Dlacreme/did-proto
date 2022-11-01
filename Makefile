SHELL = /bin/bash

SRC_FOLDER=.
NAME = did-proto
MAIN = $(SRC_FOLDER)/did.proto
CC = protoc

.PHONY: all
all: compile

.PHONY: compile
compile:
	$(CC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(MAIN)