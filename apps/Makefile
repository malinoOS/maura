parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: msh coreutils

all: msh coreutils

msh:
	@echo " GO apps/msh"
	@cd $(parentFolder)/msh; \
	go mod tidy; \
	go build -o /tmp/maura_msh -ldflags "-X main.Version=$(shell date +%y%m%d)"

coreutils: echo ls reboot ps

echo:
	@echo " GO apps/coreutils/echo"
	@cd $(parentFolder)/coreutils/echo; \
	go mod tidy; \
	go build -o /tmp/maura_echo -ldflags "-X main.Version=$(shell date +%y%m%d)"

ls:
	@echo " GO apps/coreutils/ls"
	@cd $(parentFolder)/coreutils/ls; \
	go mod tidy; \
	go build -o /tmp/maura_ls -ldflags "-X main.Version=$(shell date +%y%m%d)"

reboot:
	@echo " GO apps/coreutils/reboot"
	@cd $(parentFolder)/coreutils/reboot; \
	go mod tidy; \
	go build -o /tmp/maura_reboot -ldflags "-X main.Version=$(shell date +%y%m%d)"

ps:
	@echo " GO apps/coreutils/ps"
	@cd $(parentFolder)/coreutils/ps; \
	go mod tidy; \
	go build -o /tmp/maura_ps -ldflags "-X main.Version=$(shell date +%y%m%d)"