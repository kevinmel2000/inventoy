CURR_DIR = $(shell pwd)
MAIN_PATH = $(CURR_DIR)/main.go

update:
	@dep ensure -v

run:
	@go run $(MAIN_PATH)
