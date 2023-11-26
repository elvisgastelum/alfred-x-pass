DIST_DIR = dist
FILE_NAME = pass-x-alfred
DIST_FILE = $(DIST_DIR)/$(FILE_NAME).alfredworkflow
DIST_FILE_ZIP = $(DIST_DIR)/$(FILE_NAME).zip
BUILD_FINISHED_MESSAGE = "✅ Workflow built successfully at $(DIST_FILE)"
GO_FILTER_BUILD_DIR = build
GO_FILTER_OUTPUT_FILE = pasawutil
GO_FILTER_OUTPUT = $(GO_FILTER_BUILD_DIR)/$(GO_FILTER_OUTPUT_FILE)
GO_FILTER_SRC=src/main.go

build: $(DIST_FILE)
	@echo $(BUILD_FINISHED_MESSAGE)

$(DIST_FILE): info.plist icon.png assets/key-icon.png assets/not-found-icon.png $(GO_FILTER_OUTPUT) scripts/pass-show.sh scripts/pass-generate.sh scripts/pass-otp.sh scripts/pass-get-login-field.sh
	[ -d $(DIST_DIR) ] || mkdir -p $(DIST_DIR);
	zip $@ $^

$(GO_FILTER_OUTPUT): $(GO_FILTER_SRC)
	go build -o $@ $^

zip: info.plist $(GO_FILTER_OUTPUT) scripts/pass-show.sh scripts/pass-generate.sh scripts/pass-otp.sh
	zip $(DIST_FILE_ZIP) $^

clean:
	rm -rf dist build

dev: clean build
	open $(DIST_FILE)
