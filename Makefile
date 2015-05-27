GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOCOV=gocov
GOCOVREPORT=$(GOCOV) test | gocov-html > cover.html

BINNAME=gorender

all:
	$(GOBUILD) -o $(BINNAME)

test:
	$(GOTEST)

clean:
	$(GOCLEAN)

cover:
	$(GOCOVREPORT)
