BINNAME:=k8sr
BINDIR:=../bin
GOARCH:=amd64
GO_FILES:=../

.PHONY: build_cmd

build_cmd:
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME).exe $(GO_FILES)
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINDIR)/$(BINNAME) $(GO_FILES)
