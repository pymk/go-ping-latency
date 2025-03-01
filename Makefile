.PHONY: build install clean

BUILD_DIR = build
APP_NAME = go-ping-latency
INSTALL_DIR = /usr/local/bin

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

install: build
	install -m 0755 $(BUILD_DIR)/$(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
