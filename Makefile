.PHONY: run install

run:
	go run cmd/jn/main.go

clean:
	rm -rf bin

install:
	go build -o bin/jn cmd/jn/main.go
	mv bin/jn $$GOPATH/bin/jn
