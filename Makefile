include Makefile.include

## Add your make instructions here

PROJECT_PREFIX=github.com/lunaticthinker-me
PROJECT=demo-cross-lang-encryption-go

GIT_CERT_IGNORE =
GIT_CERT_IGNORE_COMMAND =
ifneq ($(GIT_CERT_IGNORE),)
	GIT_CERT_IGNORE_COMMAND = git config http.sslVerify false
endif

# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
BUILD_OS = linux
BUILD_ARCH = amd64
# GIT_REV_LIST = $(shell git rev-list --tags --max-count=1)
# BUILD_VERSION = $(shell git describe --tags $(GIT_REV_LIST))
BUILD_VERSION = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%%(refname)' --count=1)
ifeq ($(BUILD_VERSION),)
BUILD_VERSION=none
endif
BUILD_COMMIT = $(shell git log --format="%%h" -n 1)

BUILD_DATE = $(shell date --utc)
BUILD_PATH = ./dist/$(BUILD_OS)/$(BUILD_ARCH)
BUILD_VARS = GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH)
BUILD_SRC = ./src/main.go
BUILD_BIN = $(BUILD_PATH)/main
ifeq ($(OSFLAG),WIN32)
BUILD_DATE = $(shell $(POWERSHELL) -Command 'Get-Date -Format "yyyyMMddHHmmss"')
# BUILD_PATH = .\dist\$(BUILD_OS)\$(BUILD_ARCH)
# BUILD_VARS = set GOOS="$(BUILD_OS)"; set GOARCH="$(BUILD_ARCH)";
# BUILD_SRC = .\main.go
# BUILD_BIN = $(BUILD_PATH)\main
endif

ifeq ($(BUILD_OS),windows)
	BUILD_EXT = .exe
else
	BUILD_EXT =
endif

BUILD_VERSION_FLAG = $(PROJECT_PREFIX)/$(PROJECT).VersionName=$(BUILD_VERSION)
BUILD_COMMIT_FLAG = $(PROJECT_PREFIX)/$(PROJECT).GitCommit=$(BUILD_COMMIT)
BUILD_DATE_FLAG = $(PROJECT_PREFIX)/$(PROJECT).BuildDate=$(BUILD_DATE)

GO := GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build -trimpath
GO_LDFLAGS = -X $(BUILD_VERSION_FLAG) -X $(BUILD_COMMIT_FLAG) -X '$(BUILD_DATE_FLAG)'

#
# Instructions
#

build: test clean build-$(SHELL_IS) ## Build Application BUILD_OS=? BUILD_ARCH=? #$(BUILD_SRC)

build-run:
	$(GO) -ldflags "$(GO_LDFLAGS)" -o $(BUILD_BIN)$(BUILD_EXT) $(BUILD_SRC)

build-bash: build-bash-mkdir build-run

build-bash-mkdir:
	mkdir -p dist/$(BUILD_OS)/$(BUILD_ARCH)

build-powershell: GO = $(POWERSHELL) -File ./.scripts/make.ps1 -Action Build -Command "go build -trimpath" -GoOs $(BUILD_OS) -GoArch $(BUILD_ARCH)
build-powershell: BUILD_SRC = -Src ./src/main.go
build-powershell: build-powershell-mkdir build-run

build-powershell-mkdir:
	$(POWERSHELL) -File ./.scripts/make.ps1 -Action MkDir -Path dist\$(BUILD_OS)\$(BUILD_ARCH)

CLEAN_FULL=
clean: clean-$(SHELL_IS) ## Clean all dist/temp folders

clean-bash:
ifneq ($(CLEAN_FULL),)
	rm -rf ./dist
else
	rm -rf $(BUILD_PATH)
endif

clean-powershell:
ifneq ($(CLEAN_FULL),)
	$(POWERSHELL) -File ./.scripts/make.ps1 -Action RmDir -Path .\dist
else
	$(POWERSHELL) -File ./.scripts/make.ps1 -Action RmDir -Path $(BUILD_PATH)
endif

configure: ## Configure and Init the code dependencies
	go get -t -v ./...

	python3 -m pip install pre-commit
	python3 -m pre_commit
	python3 -m pre_commit run --all-files


install: build ## Install Application
	@echo 'Install Instructions'


RUN_ARGS  =
run: run-$(SHELL_IS) ## Run Application (from source code)

run-bash:
	go run ./src/main.go $(RUN_ARGS)

run-powershell:
	go run .\src\main.go $(RUN_ARGS)


run-binary: build run-binary-$(SHELL_IS) ## Run Application (from binary)

run-binary-bash:
	./dist/$(BUILD_OS)/$(BUILD_ARCH)/main $(RUN_ARGS)

run-binary-powershell:
	.\dist\$(BUILD_OS)\$(BUILD_ARCH)\main.exe $(RUN_ARGS)


uninstall: ## Uninstall Application
	@echo 'Uninstall Instructions'


GO_TEST_FLAGS=-tags=unit -timeout 30s -short -coverprofile=.coverage/coverage -v
GO_TEST=go test $(GO_TEST_FLAGS)
GO_COVERAGE_FLAGS=-o .coverage/coverage.html -html=.coverage/coverage
GO_COVERAGE=go tool cover $(GO_COVERAGE_FLAGS)
test: test-$(SHELL_IS) ## Run Tests
	$(GO_COVERAGE)

# test-bash:
# 	$(GO_TEST) ./...

test-bash:
	mkdir -p ./.coverage
	find ./src -iname "*_test.go" | while read f; do echo $$(dirname $$f)/...; done | uniq | xargs $(GO_TEST)

# test-powershell:
# 	$(GO_TEST) .\...

test-powershell:
	$(POWERSHELL) -File ./.scripts/make.ps1 -Action MkDir -Path .\.coverage
	$(POWERSHELL) -File ./.scripts/make.ps1 -Action Test -Command "$(GO_TEST)"


TEST_PATH=./...
test-single:
	$(GO_TEST) $(TEST_PATH)
