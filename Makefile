all: fmt test
	@mkdir -p bin/
	@bash --norc -i ./scripts/build.sh

fmt:
	@echo "\n==> Formatting source code\n"
	@go fmt ./...

test:
	go list ./... | xargs -n1 go test
