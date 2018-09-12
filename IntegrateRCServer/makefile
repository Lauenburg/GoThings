#Makefile for IntegrateRCServer

BUILDPATH=$(CURDIR)
EXENAME=OTemperature

#export GOPATH=$(CURDIR)

help:
	@echo "Make file for IntegrateRCServer. Type 'make all' to build tree, load packages and build."
	@echo "This make file expects that vgo installed (go get -u golang.org/x/vgo)"

init:
	$(shell vgo mod init)

deps:
	@echo "getting the packages..."
	$(shell vgo get github.com/TheThingsNetwork/go-app-sdk)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log/apex)
	$(shell vgo get github.com/Lauenburg/GoThings/IntegrateRCServer/oTemp)
	$(shell vgo get github.com/gorilla/mux)


build:
	@echo "start building..."
	$(shell vgo build)

all: init deps build
