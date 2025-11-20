.DEFAULT_GOAL := help

BIN_NAME := evm-wallet
BUILD_DIR := ./out

##@ Build

.PHONY: $(BUILD_DIR)
$(BUILD_DIR): ## Create the build folder.
	mkdir -p $(BUILD_DIR)

.PHONY: build
build: $(BUILD_DIR) ## Build go binary.
	go build -o $(BUILD_DIR)/$(BIN_NAME) main.go

.PHONY: cross
cross: $(BUILD_DIR) ## Cross-compile go binaries without using CGO.
	mkdir -p $(BUILD_DIR)/$(BIN_NAME)_linux_amd64
	mkdir -p $(BUILD_DIR)/$(BIN_NAME)_darwin_amd64
	GOOS=linux  GOARCH=amd64 go build -o $(BUILD_DIR)/$(BIN_NAME)_linux_amd64/$(BIN_NAME)  main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BIN_NAME)_darwin_amd64/$(BIN_NAME) main.go


.PHONY: clean
clean: ## Clean the binary folder.
	$(RM) -r $(BUILD_DIR)
