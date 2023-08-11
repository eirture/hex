
.PHONY: build
build:
	go build .

.PHONY: clean
clean:
	rm hex

.PHONY: install
install:
	go install .

.PHONY: uninstall
uninstall:
	rm $(GOPATH)/bin/hex