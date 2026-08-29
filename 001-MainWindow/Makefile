# Makefile for Go + walk

APP := $(notdir $(CURDIR))
	
all: $(APP).exe

$(APP).exe: $(APP).go go.mod resource.syso
	go build -trimpath -ldflags="-s -w -H windowsgui" $(APP)

run: $(APP).go go.mod resource.syso
	go run $(APP)

go.mod: $(APP).go
	rm -f go.mod
	go mod init $(APP)
	go mod tidy
	go get github.com/lxn/walk

resource.syso: versioninfo.json ../files/gowalk.manifest ../files/icon.ico
	go generate

clean:
	rm -f $(APP).exe resource.syso go.mod go.sum
