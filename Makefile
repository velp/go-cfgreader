GOPATH = $(shell echo $$GOPATH)
BUILD_DIR = _output

clean:
	rm -rf $(BUILD_DIR)

deps:
	go get ./...

build: $(BUILD_DIR)
	go build -o $(BUILD_DIR)/cfgreader

test: test-deps
	$(GOPATH)/bin/ginkgo -r --randomizeAllSpecs -p -nodes=4

test-deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)