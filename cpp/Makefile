
PROC_COUNT ?= 8

COLOR_HIGHLIGHT_BLUE=\e[1;36m
COLOR_HIGHLIGHT_GREEN=\e[1;32m
COLOR_CLEAR=\e[0m

all: setup install

setup:
	@mkdir -p bin

install: setup install-libs release
	@echo "$(COLOR_HIGHLIGHT_BLUE)Installing toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make install

install-libs: toy-lib

%-lib:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Building $*-lib...$(COLOR_CLEAR)"
	cd src/$*-lib && make install

release:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Building toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make release

clean-libs: clean-toy-lib

clean-%-lib:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning $*-lib...$(COLOR_CLEAR)"
	cd src/$*-lib && make clean
	rm -rf lib/lib$**
	rm -rf include/$*-lib

clean: clean-libs
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make clean
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning binaries...$(COLOR_CLEAR)"
	rm -rf bin/* include/*
