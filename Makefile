NAME			=	go-api

BUILD_NAME		=	$(NAME)

BUILD_DIR		=	./build

CMD_DIR			=	./cmd

GO				=	go

GO_BUILD		=	$(GO) build

GO_MOD			=	$(GO) mod

GO_GET			=	$(GO) get

GO_CLEAN		=	$(GO) clean

GO_TEST			=	$(GO) test

GO_ENV			=	$(GO) env

RM				=	rm -f

MKDIR			=	mkdir -p

FLAGS			=	-p 2

all				:	build

build			:
					$(MKDIR) $(BUILD_DIR)
					$(GO_BUILD) -o $(BUILD_DIR)/$(BUILD_NAME) -v $(FLAGS) $(CMD_DIR)/$(BUILD_NAME)/main.go

test			:
					$(GO_TEST) -v ./...

vet 			:
					$(GO) vet ./...

clean			:
					$(GO_CLEAN)
					$(RM) $(BUILD_DIR)/$(BUILD_NAME)

install			:
					$(GO_MOD) download

update			:
					$(GO_GET) -u ./...
					$(GO_MOD) tidy
					$(GO_MOD) vendor

vendor			:
					$(GO_MOD) vendor

rebuild			:	clean build

re				:	rebuild

.PHONY			:	all make build test vet clean install update vendor rebuild re
