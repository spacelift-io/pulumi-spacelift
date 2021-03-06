PROJECT_NAME := pulumi-spacelift Package

PACK             := spacelift
ORG              := spacelift-io
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @pulumi/${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION_WITH_V  := "$(shell git describe --tags --abbrev=0)"
VERSION         := "$(shell git describe --tags --abbrev=0 | sed 's/^v//g')"
PROVIDER_OS     := ${PROVIDER_OS}

TESTPARALLELISM := 4

WORKING_DIR     := $(shell pwd)

OS := $(shell uname)
EMPTY_TO_AVOID_SED := ""

prepare::
	@if test -z "${NAME}"; then echo "NAME not set"; exit 1; fi
	@if test -z "${REPOSITORY}"; then echo "REPOSITORY not set"; exit 1; fi
	@if test ! -d "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz"; then "Project already prepared"; exit 1; fi

	mv "provider/cmd/pulumi-tfgen-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-tfgen-${NAME}
	mv "provider/cmd/pulumi-resource-x${EMPTY_TO_AVOID_SED}yz" provider/cmd/pulumi-resource-${NAME}

	if [[ "${OS}" != "Darwin" ]]; then \
		sed -i 's,github.com/pulumi/pulumi-xyz,${REPOSITORY},g' provider/go.mod; \
		find ./ ! -path './.git/*' -type f -exec sed -i 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
	fi

	# In MacOS the -i parameter needs an empty string to execute in place.
	if [[ "${OS}" == "Darwin" ]]; then \
		sed -i '' 's,github.com/pulumi/pulumi-xyz,${REPOSITORY},g' provider/go.mod; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \

		sed -i '' 's,github.com/pulumi/pulumi-xyz,${REPOSITORY},g' provider/resources.go; \
		find ./ ! -path './.git/*' -type f -exec sed -i '' 's/[x]yz/${NAME}/g' {} \; &> /dev/null; \
	fi

.PHONY: development provider build_sdks build_nodejs build_dotnet build_go build_python cleanup

development:: install_plugins provider lint_provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -a -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && CGO_ENABLED=0 GOOS=$(PROVIDER_OS) go build -a -tags netgo -o $(WORKING_DIR)/bin/ -ldflags "-s -w -extldflags '-static' -X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_python build_go build_dotnet # build all the sdks

build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
    	sed -i.bak -e "s/\$${VERSION}/$(VERSION_WITH_V)/g" ./bin/package.json

build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	if [[ "${OS}" == "Darwin" ]]; then \
        sed -i '' -r "s/check_call\(\['pulumi', 'plugin', 'install', 'resource', 'spacelift', '[^']+']\)/check_call(['pulumi', 'plugin', 'install', '--server', 'https:\/\/downloads.spacelift.io\/pulumi-plugins', 'resource', 'spacelift', '0.0.0'])/g" sdk/python/setup.py; \
    fi
	if [[ "${OS}" != "Darwin" ]]; then \
        sed -i -r "s/check_call\(\['pulumi', 'plugin', 'install', 'resource', 'spacelift', '[^']+']\)/check_call(['pulumi', 'plugin', 'install', '--server', 'https:\/\/downloads.spacelift.io\/pulumi-plugins', 'resource', 'spacelift', '0.0.0'])/g" sdk/python/setup.py; \
    fi
	if [[ "${OS}" == "Darwin" ]]; then \
        sed -i '' -r "s/import pulumi$$/import pulumi as pulumilib/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/pulumi.CustomResource/pulumilib.CustomResource/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/pulumi.ResourceOptions/pulumilib.ResourceOptions/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/pulumi.Input/pulumilib.Input/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/pulumi.Output/pulumilib.Output/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/@pulumi.getter/@pulumilib.getter/g" sdk/python/pulumi_spacelift/stack.py && \
        sed -i '' -r "s/pulumi.get/pulumilib.get/g" sdk/python/pulumi_spacelift/stack.py; \
    fi
	if [[ "${OS}" != "Darwin" ]]; then \
        sed -i -r "s/import pulumi$$/import pulumi as pulumilib/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/pulumi.CustomResource/pulumilib.CustomResource/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/pulumi.ResourceOptions/pulumilib.ResourceOptions/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/pulumi.Input/pulumilib.Input/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/pulumi.Output/pulumilib.Output/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/@pulumi.getter/@pulumilib.getter/g" sdk/python/pulumi_spacelift/stack.py && \
		sed -i -r "s/pulumi.get/pulumilib.get/g" sdk/python/pulumi_spacelift/stack.py; \
    fi
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/0.0.0/g" ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

build_dotnet:: install_plugins tfgen # build the dotnet sdk
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "$(VERSION)" >version.txt && \
        cp ../../Pulumi.Spacelift.csproj.template ./Pulumi.Spacelift.csproj && \
        dotnet build --configuration Release /p:Version=$(VERSION)

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::
	rm -rf sdk/{dotnet,nodejs,go,python}

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource random 2.2.0

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

test::
	cd examples && go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h
