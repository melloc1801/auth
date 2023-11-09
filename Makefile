include local.env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/auth_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth_v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DR} postgres ${PG_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DR} postgres ${PG_DSN} down -v


copy-to-server:
	scp bin/main root@188.124.39.8:

docker-run:
	docker run -p 50051:50051 cr.selcloud.ru/microservices/test-server:v0.0.1

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/microservices/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAAhhJUhYQ-SgggqO7Ezr4EK31kQvqpes34 cr.selcloud.ru/microservices
	docker push cr.selcloud.ru/microservices/test-server:v0.0.1

#docker-build-and-push:
#	docker buildx build --no-cache --platform linux/amd64 -t <REGESTRY>/test-server:v0.0.1 .
#	docker login -u <USERNAME> -p <PASSWORD> <REGESTRY>
#	docker push <REGESTRY>/test-server:v0.0.1