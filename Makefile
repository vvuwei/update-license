-include .env
export $(shell sed 's/=.*//' .env)

init:
	chmod +x init.dev.sh && ./init.dev.sh

clean:
	rm -rf bin
	rm -rf update-license

lint:
	golangci-lint run

test:
	go test -timeout 5s .

bench:
	go test -bench=.

cover:
	go test . -coverpkg=. -coverprofile ./coverage.out
	go tool cover -func ./coverage.out

dry:
	go run . -path=./tmp -license=./tmp/LICENSE -dry

.PHONY: init clean lint test cover dry bench
