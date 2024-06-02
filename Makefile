parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: all clean msh run iso

all: clean msh run

clean:
	-rm /tmp/maura_msh

msh:
	cd $(parentFolder)/apps/msh; \
	go mod tidy; \
	go build -o /tmp/maura_msh -ldflags "-X main.Version=$(shell date +%y%m%d)"

iso:
	malino build
	malino export -efi

run:
	malino build
	malino run