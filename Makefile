parentFolder := $(shell pwd)
SHELL := /bin/bash

.PHONY: all clean run iso apps prepare

all: clean os-release apps run clean

clean:
	@-if [ -a os-release.maura ] ; \
	then \
		echo " RM *.maura" ; \
		rm *.maura ; \
	fi;
	@-rm maura.iso initramfs.cpio.gz

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
	@echo "  W os-release.maura"
	@printf "PRETTY_NAME=\"maura $(shell date +%y%m%d)\"\nNAME=\"maura malino/Linux\"\nVERSION_ID=\"$(shell date +%y%m%d)\"\nVERSION=\"v$(shell date +%y%m%d)\"\nVERSION_CODENAME=v$(shell date +%y%m%d)\nID=maura" > os-release.maura

apps:
ifeq (,$(wildcard ./fastfetch))
    $(error please run make prepare first)
endif
	make -C apps/

iso:
	malino build
	malino export

run:
	malino build
	malino run