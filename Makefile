.PHONY: all clean
# 被编译的文件
BUILDFILE=*.go
# 编译后的静态链接库文件名称
TARGETNAME=crddemo
# GOOS 为目标主机系统
GOOS=linux
GOARCH=amd64

all: format test build clean

test:
	go test -v .

format:
	gofmt -w .

build:
	mkdir -p releases
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod=vendor -ldflags "-s -w" -v -o releases/$(TARGETNAME) $(BUILDFILE)

clean:
	go clean -i