#Makefile for OTemperature

BUILDPATH=$(CURDIR)
VGO=$(shell which vgo)
VGOINIT = $(VGO) mod init
VGOBUILD=$(VGO) build
VGOCLEAN=$(VGO) clean
VGOGET=$(VGO) get

EXENAME=OTemperature

export GOPATH=$(CURDIR)
export GOBIN=$(CURDIR/bin)

myname:
	@echo "Make file for OTemperature. Type 'make all' to build tree, load packages and build."

init:
	$(shell vgo mod init)
deps:
	@echo "getting the packages..."
	$(shell vgo get github.com/TheThingsNetwork/go-app-sdk)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log)
	$(shell vgo get github.com/TheThingsNetwork/go-utils/log/apex)
	@echo "all done"

build:
	@echo "start building..."
	$(shell vgo build)
	@echo "build was successful"

clean:
	@echo "cleanning"
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg
	@rm -rf $(BUILDPATH)/src/github.com

all: init deps build
