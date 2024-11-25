VERSION			?= $(shell pulumictl get version)

PACK			:= svmkit
PROJECT			:= github.com/abklabs/pulumi-${PACK}

PROVIDER		:= pulumi-resource-${PACK}

GOLANG_FLAGS		= -ldflags "-X ${PROJECT}/pkg/version.Version=${VERSION}"
GOPATH			:= $(shell go env GOPATH)
TESTPARALLELISM		:= 4
GO_TEST			:= go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

WORKING_DIR		:= $(shell pwd)
EXAMPLES_DIR		:= ${WORKING_DIR}/examples/yaml

OS			:= $(shell uname)

all: build install

ensure::
	go work sync
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

gen_examples: gen_go_example \
		gen_nodejs_example \
		gen_python_example \
		gen_dotnet_example

gen_%_example:
	rm -rf ${WORKING_DIR}/examples/$*
	pulumi convert \
		--cwd ${WORKING_DIR}/examples/yaml \
		--logtostderr \
		--generate-only \
		--non-interactive \
		--language $* \
		--out ${WORKING_DIR}/examples/$*

.PHONY: build

build:: provider dotnet_sdk go_sdk nodejs_sdk python_sdk

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

test_all:: test_provider
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/dotnet && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/AbkLabs.Svmkit.*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

clean::
	$(RM) -r bin nuget sdk/dotnet/{bin,obj} sdk/nodejs/bin sdk/python/bin

distclean:: clean
	$(RM) -r sdk/nodejs/node_modules
