PROJECT_ID=lisyaoran51
OS=$(shell uname | tr '[:upper:]' '[:lower:]')
SERVICE_NAME=$(shell basename `git rev-parse --show-toplevel`)
IMPORT_PATH=github.com/lisyaoran51/${SERVICE_NAME}
PROTOS=$(shell ls ./rf-protos/models/)
GIT_COMMIT_HASH=$(shell git rev-parse HEAD | cut -c -16)
BUILD_TIME=$(shell date +%s)
LDFLAGS = -X ${IMPORT_PATH}/global.ServiceName=${SERVICE_NAME}
LDFLAGS += -X ${IMPORT_PATH}/global.GitCommitHash=${GIT_COMMIT_HASH}
LDFLAGS += -X ${IMPORT_PATH}/global.BuildTime=${BUILD_TIME}
TAG=${PROJECT_ID}/${SERVICE_NAME}:${GIT_COMMIT_HASH}
IMAGE=lisyaoran51/${TAG}
ifeq (${OS}, darwin)
	SED_INPLACE = sed -i'.orig' -e
endif
ifeq (${OS}, linux)
	SED_INPLACE = sed -i
endif

.PHONY: proto proto2

proto:
	@for file in `ls proto/models/`; do \
		protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false  --proto_path=./models ./models/$${file}/*.proto; \
	done ;
	find ./models/* -name '*.pb.go' -exec ${SED_INPLACE} s/,omitempty//g '{}' \; ;
	find ./models -name '*.orig' -delete ;

proto2:
	@for file in `ls ./proto/`; do \
		protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false  --proto_path=./proto ./proto/$${file}/*.proto; \
	done ;
	cp -r github.com/paper-trade-chatbot/be-proto/* ./
	rm -r github.com