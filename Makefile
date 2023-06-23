GO := go
SRC := .
BIN := rgn
BIN_DIR := ./bin
BIN_LOC := $(BIN_DIR)/$(BIN)
GOPATH := ${GOPATH}

.PHONY: build install clean test

build:
	@mkdir -p ./bin
	$(GO) build -o $(BIN_LOC) $(SRC)

test:
	$(GO) test -v ./...

install:build
	@mv -v $(BIN_LOC) $(GOPATH)/bin

clean:
	@rm -r $(BIN_DIR)
	@echo "Removed `./bin` directory"


# Remove rgn from sys
uninstall:clean
	@rm -v $(GOPATH)/bin/$(BIN)
