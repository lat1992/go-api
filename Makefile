NAME			=	enigm

GO				=	go

GOBUILD			=	$(GO) build

GOCLEAN			=	$(GO) clean

GOTEST			=	$(GO) test

GOGET			=	$(GO) get

GOENV			=	$(GO) env

RM				=	rm -f

FLAGS			=	-p 2

DEBUG_FLAGS		=	-gcflags '-N -l'

all				:	test build

build			:
					$(GOBUILD) -o $(NAME) -v $(FLAGS)

debug_build		:
					$(GOBUILD) -o $(NAME) -v $(DEBUG_FLAGS)

test			:
					$(GOTEST) -v

run				:
					./$(NAME)

clean			:
					$(GOCLEAN)
					$(RM) $(NAME)

dependencies	:
					$(GOGET) -u golang.org/x/text
					$(GOGET) -u github.com/gin-gonic/gin
					$(GOGET) -u github.com/gin-contrib/cors
					$(GOGET) -u github.com/go-pg/pg
					$(GOGET) -u github.com/dchest/uniuri

install			:	dependencies

rebuild			:	clean test build

re				:	rebuild

debug_rebuild	:	clean test debug_build

dre				:	debug_rebuild

.PHONY			:	all make build debug_build test clean run dependencies install rebuild re
