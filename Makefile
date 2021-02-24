NAME			=	go-api

BINARY_NAME		=	$(NAME).out

GO				=	go

GOBUILD			=	$(GO) build

GOMOD			=	$(GO) mod

GOCLEAN			=	$(GO) clean

GOTEST			=	$(GO) test

GOENV			=	$(GO) env

RM				=	rm -f

FLAGS			=	-p 2

all				:	build

build			:
					$(GOBUILD) -o $(BINARY_NAME) -v $(FLAGS)

test			:
					$(GOTEST) -v ./...

run				:
					./$(BINARY_NAME)

clean			:
					$(GOCLEAN)
					$(RM) $(BINARY_NAME)

install			:
					$(GOMOD) download

vendor			:
					$(GOMOD) vendor

rebuild			:	clean build

re				:	rebuild

.PHONY			:	all make build test clean run install vendor rebuild re
