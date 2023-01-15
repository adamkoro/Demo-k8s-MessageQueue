build:
	go build -o server main.go

run: build
	./server

tidy:
	go mod tidy

watch:
	~/go/bin/reflex -s -r '\.go$$' make run