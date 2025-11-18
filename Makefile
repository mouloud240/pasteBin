install:
	go mod tidy
compile:
	go build -o bin/pasteBin cmd/api/main.go
run:
	bin/pastebin
gen_doc:
	swag init -g cmd/api/main.go -o cmd/api/docs
build: install compile
all: install compile run
