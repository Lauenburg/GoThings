#Makefile for RCServer
BUILDPATH=$(CURDIR)
EXENAME=RCServer

#export GOPATH=$(CURDIR)

help:
	@echo "Make file for RCServer. Type 'make all' to load packages and build."
	@echo "This make file expects that vgo installed (go get -u golang.org/x/vgo)"
	@echo "Write 'make all' to build"

init:
	$(shell vgo mod init)

deps:
	@echo "getting the packages..."
	$(shell vgo get github.com/gorilla/mux)

build:
	@echo "start building..."
	$(shell vgo build)

all: init build
