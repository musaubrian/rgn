GO := go
SRC := .
BIN := rgn
BIN_LOC := ./bin/$(BIN)
GOPATH := ${GOPATH}

.PHONY: build install

build:
	@mkdir -p ./bin
	$(GO) build -o $(BIN_LOC) $(SRC)

install:build
	@mv -v $(BIN_LOC) $(GOPATH)/bin
