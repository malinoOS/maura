parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: msh coreutils

all: msh coreutils

msh:
	@echo " GO apps/msh"
	@cd $(parentFolder)/msh; \
	go mod tidy; \
	go build -o /tmp/maura_msh -ldflags "-X main.Version=$(shell date +%y%m%d)"

coreutils: cat echo ls reboot ps

cat:
	@echo " GO apps/coreutils/cat"
	@cd $(parentFolder)/coreutils/cat; \
	go mod tidy; \
	go build -o /tmp/maura_cat -ldflags "-X main.Version=$(shell date +%y%m%d)"

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