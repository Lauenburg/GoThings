#Makefile for OTemperature

BUILDPATH=$(CURDIR)
VGO=$(shell which vgo)
VGOINIT = $(VGO) mod init
VGOBUILD=$(VGO) build
VGOCLEAN=$(VGO) clean
VGOGET=$(VGO) get

#EXENAME=OTemperature

#export GOPATH=$(CURDIR)

help:
	@echo "Make file for OTemperature. Type 'make all' to build tree, load packages and build."
	@echo "This make file expects that vgo is installed (go get -u golang.org/x/vgo)"
	@echo "Write 'make all' to build"
init:
	$(shell vgo mod init)
deps:
	@echo "getting the packages..."
	$(shell vgo get github.com/TheThingsNetwork/go-app-sdk)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log/apex)

build:
	@echo "start building..."
	$(shell vgo build)

all: init deps build
