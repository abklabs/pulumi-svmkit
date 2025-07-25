VERSION			?= $(shell pulumictl get version)

PACK			:= svmkit
PROJECT			:= github.com/abklabs/pulumi-${PACK}

PROVIDER		:= pulumi-resource-${PACK}

GOLANG_FLAGS		= -ldflags "-X ${PROJECT}/pkg/version.Version=${VERSION}"
GOPATH			:= $(shell go env GOPATH)
TESTPARALLELISM		:= 4
GO_TEST			:= go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

WORKING_DIR		:= $(shell pwd)

OS			:= $(shell uname)

all: build install

ensure::
	cd provider && go mod tidy
	cd sdk/go && go mod tidy

provider $(WORKING_DIR)/bin/$(PROVIDER)::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} ${GOLANG_FLAGS} $(PROJECT)/cmd/$(PROVIDER))

provider_debug:: GOLANG_FLAGS += -gcflags="all=-N -l"
provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} ${GOLANG_FLAGS} $(PROJECT)/cmd/$(PROVIDER))

dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
dotnet_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/dotnet
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language dotnet
	cd sdk/dotnet/&& \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

go_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/go
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language go
	cd sdk/go && go mod init $(PROJECT)/sdk/go && go mod tidy && go mod edit -go=1.22
	$(MAKE) ensure

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/nodejs
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language nodejs
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock bin/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' bin/package.json && \
		rm ./bin/package.json.bak

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/python
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language python
	cp README.md sdk/python/
	cd sdk/python/ && \
		python3 setup.py clean --all && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

.PHONY: build

build:: provider dotnet_sdk go_sdk nodejs_sdk python_sdk

install:: install_nodejs_sdk install_dotnet_sdk

test:: provider
	cd provider && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/AbkLabs.Svmkit.*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

clean::
	$(RM) -r bin/pulumi-resource-svmkit nuget sdk/dotnet/{bin,obj} sdk/nodejs/bin sdk/python/bin

distclean:: clean
	$(RM) -r sdk/nodejs/node_modules

lint:
	(cd provider &&  golangci-lint run --path-prefix ./provider )
	shfmt -d .githooks/*
	@diff=$$(gofmt -d provider); \
	if [ -n "$$diff" ]; then \
		echo "$$diff"; \
		echo "Formatting issues found in 'provider'. Run 'gofmt -w provider' to fix."; \
		exit 1; \
	fi
	shellcheck -P .githooks .githooks/*

check: test lint

format:
	(cd provider && go fmt ./...)
	shfmt -w .githooks/*
