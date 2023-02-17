
PROC_COUNT ?= 8

COLOR_HIGHLIGHT_BLUE=\e[1;36m
COLOR_HIGHLIGHT_GREEN=\e[1;32m
COLOR_CLEAR=\e[0m

install: setup release
	@echo "$(COLOR_HIGHLIGHT_BLUE)Installing toy-lib...$(COLOR_CLEAR)"
	cd src/toy-lib && make install
	@echo "$(COLOR_HIGHLIGHT_BLUE)Installing toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make install

setup:
	@mkdir -p bin

release:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Building toy-lib...$(COLOR_CLEAR)"
	cd src/toy-lib && make release
	@echo "$(COLOR_HIGHLIGHT_BLUE)Building toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make release

clean:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning toy-lib...$(COLOR_CLEAR)"
	cd src/toy-lib && make clean
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning toy-app...$(COLOR_CLEAR)"
	cd src/toy-app && make clean
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning binaries...$(COLOR_CLEAR)"
	rm -rf lib/* bin/* include/*
