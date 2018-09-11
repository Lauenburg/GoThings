#Makefile for RCServer

BUILDPATH=$(CURDIR)
EXENAME=RCServer

export GOPATH=$(CURDIR)
export GOBIN=$(CURDIR/bin)

myname:
	@echo "Make file for RCServer. Type 'make all' to build tree, load packages and build."
	@echo "This make file expects that vgo installed (go get -u golang.org/x/vgo)"
init:
	$(shell vgo mod init)

build:
	@echo "start building..."
	$(shell vgo build)

all: init build
