.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	golint

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables

# All gen targets
.PHONY: gen
gen: gen-docs