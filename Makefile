export GO111MODULE=on
export GOSUMDB=off

BUILD_ENVPARMS:=CGO_ENABLED=0

LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_FOLDER=$(CURDIR)/scripts/migrations
PROTOGEN_BIN:=$(LOCAL_BIN)/protogen
MINIMOCK_BIN:=$(LOCAL_BIN)/minimock
MINIMOCK_TAG:=3.0.10
LINTER_TAG:=1.51.2
GOOSE_TAG:=latest

ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=kitchen-test host=localhost port=6432 sslmode=disable
endif

# set if not set
ifeq ($(GOOSE_BIN),)
	GOOSE_BIN := $(LOCAL_BIN)/goose
endif

# install goose
.PHONY: install-goose
install-goose:
ifeq ($(wildcard $(GOOSE_BIN)),)
	$(info "#Downloading goose $(GOOSE_TAG)")
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/pressly/goose/v3/cmd/goose@$(GOOSE_TAG) && \
		go build -o $(GOOSE_BIN) github.com/pressly/goose/v3/cmd/goose && \
		rm -rf $$tmp
endif

# install minimock binary
.PHONY: install-minimock
install-minimock:
ifeq ($(wildcard $(MINIMOCK_BIN)),)
	echo "#Downloading minimock v$(MINIMOCK_TAG)"
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/gojuno/minimock/v3/cmd/minimock@v$(MINIMOCK_TAG) && \
		go build -ldflags "-X 'main.version=$(MINIMOCK_TAG)' -X 'main.commit=test' -X 'main.buildDate=test'" -o $(LOCAL_BIN)/minimock github.com/gojuno/minimock/v3/cmd/minimock && \
		rm -rf $$tmp
endif

.PHONY: install-lint
install-lint: ## install golangci-lint binary
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(LINTER_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(LINTER_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: lint
lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.pipeline.yaml ./...

.PHONY: bin-deps
bin-deps:
	$(info #Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.16.0

# docker-compose aliases
.PHONY: compose-up
compose-up:
	docker-compose -p kitchen -f ./docker-compose.yml up -d

.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up

.PHONY: compose-rm
compose-rm:
	docker-compose -p kitchen -f /docker-compose.yml rm -fvs

.PHONY: compose-down
compose-down:
	docker-compose -p kitchen -f /docker-compose.yml stop

.PHONY: generate
generate:
	protoc \
		-I vendor.protogen \
		--plugin=bin/protoc-gen-grpc \
		--plugin=bin/protoc-gen-grpc-gateway \
		--plugin=bin/protoc-gen-swagger \
		--plugin=bin/protoc-gen-go \
	    --proto_path=api/kitchen_api \
		--proto_path=vendor.protogen \
		--go_out=pkg/kitchen_api \
		--go-grpc_out=pkg/kitchen_api \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt generate_unbound_methods=true \
		--grpc-gateway_opt path=pkg/kitchen_api \
		--swagger_out=logtostderr=true:pkg/kitchen_api \
		api/kitchen_api/*.proto

# create new migration
.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migrations-up
test-migrations-up:
	$(BUILD_ENVPARMS) $(GOOSE_BIN) -dir "$(MIGRATION_FOLDER)" postgres "${POSTGRES_SETUP_TEST}" up

.PHONY: test-migrations-down
test-migrations-down:
	$(BUILD_ENVPARMS) $(GOOSE_BIN) -dir "$(MIGRATION_FOLDER)" postgres "${POSTGRES_SETUP_TEST}" down
