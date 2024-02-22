.PHOBY: build
build:
	go build -ldflags "-w -s"

.PHOBY: build linux
build linux:
	GOOS="linux" GOARCH="amd64" go build main.go

.PHOBY: run
run:
	go run main.go

.DEFAULT_GOAL := run