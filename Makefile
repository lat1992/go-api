NAME			=	go-api

BINARY_NAME		=	$(NAME).out

GO				=	go

GOBUILD			=	$(GO) build

GOCLEAN			=	$(GO) clean

GOTEST			=	$(GO) test

GOGET			=	$(GO) get

GOENV			=	$(GO) env

RM				=	rm -f

FLAGS			=	-p 2

all				:	test build

build			:
					$(GOBUILD) -o $(BINARY_NAME) -v $(FLAGS)

test			:
					$(GOTEST) -v ./...

run				:
					./$(BINARY_NAME)

clean			:
					$(GOCLEAN)
					$(RM) $(BINARY_NAME)

dependencies	:
					$(GOGET) -u github.com/gin-gonic/gin
					$(GOGET) -u github.com/gin-contrib/cors
					$(GOGET) -u github.com/go-pg/pg
					$(GOGET) -u github.com/dchest/uniuri

install			:	dependencies

rebuild			:	clean build

re				:	rebuild

.PHONY			:	all make build test clean run dependencies install rebuild re
