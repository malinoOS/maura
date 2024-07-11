parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: msh coreutils

all: msh coreutils

msh:
	@echo " GO apps/msh"
	@cd $(parentFolder)/msh; \
	go mod tidy; \
	go get -u github.com/malinoOS/malino/libmalino; \
	go build -o $(parentFolder)/../msh.maura -ldflags "-X main.Version=$(shell date +%y%m%d)"

coreutils:
	@for dir in cat echo ls reboot ps; do \
		echo " GO apps/coreutils/$$dir"; \
		cd $(parentFolder)/coreutils/$$dir; \
		go mod tidy; \
		go get -u github.com/malinoOS/malino/libmalino; \
		go build -o $(parentFolder)/../$$dir.maura -ldflags "-X main.Version=$(shell date +%y%m%d)"; \
	done
