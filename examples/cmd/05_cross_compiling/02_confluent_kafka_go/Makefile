BUILD_DIR = target/build
RD_KAFKA_BUILD_TAG ?= static

# -----

default: run


all: build


build: clean
	mkdir -p $(BUILD_DIR)/bin
	mkdir -p $(BUILD_DIR)/lib

	go build -ldflags="-s -w" -o $(BUILD_DIR)/bin/test ./cmd
	go build -ldflags="-s -w" -tags $(RD_KAFKA_BUILD_TAG) -o $(BUILD_DIR)/lib/libTest.so -buildmode=c-shared ./lib/interface

clean:
	rm -fr $(BUILD_DIR)

run: build
	cd $(BUILD_DIR) && ./bin/test
