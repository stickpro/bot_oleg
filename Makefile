.PHONY:
.SILENT:
.DEFAULT_GOAL := run

OUT_BIN ?= ./.bin/bot-oleg
GO_LDFLAGS ?=
GO_OPT_BASE := -ldflags "-X main.version=$(VERSION) $(GO_LDFLAGS) -X main.commitHash=$(COMMIT_HASH)"

BUILD_ENV := CGO_ENABLED=0
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Linux)
	BUILD_ENV += GOOS=linux
endif
ifeq ($(UNAME_S), Darwin)
	BUILD_ENV += GOOS=darwin
endif

UNAME_P := $(shell uname -p)
ifeq ($(UNAME_P),x86_64)
	BUILD_ENV += GOARCH=amd64
endif
ifneq ($(filter arm%,$(UNAME_P)),)
	BUILD_ENV += GOARCH=arm64
endif

build:
	go mod download && $(BUILD_ENV) && go build $(GO_OPT_BASE) -o $(OUT_BIN) ./cmd/app

run: build
	$(OUT_BIN) $(filter-out $@,$(MAKECMDGOALS))