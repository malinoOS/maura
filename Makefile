parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: all clean run iso apps

all: clean os-release apps run

clean:
	@-if [ -a /tmp/maura_ps ] ; \
	then \
		echo " RM /tmp/maura_msh" ; \
		rm /tmp/maura_msh ; \
		echo " RM /tmp/maura_echo" ; \
		rm /tmp/maura_echo ; \
		echo " RM /tmp/maura_ls" ; \
		rm /tmp/maura_ls ; \
		echo " RM /tmp/maura_reboot" ; \
		rm /tmp/maura_reboot ; \
		echo " RM /tmp/maura_ps" ; \
		rm /tmp/maura_ps ; \
	fi;

prepare:
	@echo " DL https://github.com/fastfetch-cli/fastfetch/releases/download/2.14.0/fastfetch-linux-amd64.tar.gz"
	@curl -L https://github.com/fastfetch-cli/fastfetch/releases/download/2.14.0/fastfetch-linux-amd64.tar.gz --output fastfetch-linux-amd64.tar.gz -s
	@echo " GZ fastfetch-linux-amd64.tar.gz"
	@tar -xzf fastfetch-linux-amd64.tar.gz
	@echo " MV fastfetch-linux-amd64/usr/bin/fastfetch fastfetch"
	@mv fastfetch-linux-amd64/usr/bin/fastfetch fastfetch
	@echo " RM fastfetch-linux-amd64.tar.gz"
	@rm fastfetch-linux-amd64.tar.gz
	@echo " RM fastfetch-linux-amd64"
	@rm -r fastfetch-linux-amd64

os-release:
	@echo "  W /tmp/maura_os-release"
	@printf "PRETTY_NAME=\"maura $(shell date +%y%m%d)\"\nNAME=\"maura malino/Linux\"\nVERSION_ID=\"$(shell date +%y%m%d)\"\nVERSION=\"v$(shell date +%y%m%d)\"\nVERSION_CODENAME=v$(shell date +%y%m%d)\nID=maura" > /tmp/maura_os-release

apps:
ifeq (,$(wildcard ./fastfetch))
    $(error please run make prepare first)
endif
	cp ./fastfetch /tmp/maura_fastfetch
	make -C apps/

iso:
	malino build
	malino export -efi

run:
	malino build
	malino run