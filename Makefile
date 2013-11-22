all: test
	@mkdir -p bin/
	@bash --norc -i ./scripts/build.sh

test:
	go list ./... | xargs -n1 go test
