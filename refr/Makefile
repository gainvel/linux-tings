BINARY  := refr
SRC     := ./cmd/refr
INSTALL := $(HOME)/.local/bin

.PHONY: build install test lint clean

build:
	go build -o $(BINARY) $(SRC)

install: build
	mkdir -p $(INSTALL)
	cp $(BINARY) $(INSTALL)/$(BINARY)

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BINARY)
