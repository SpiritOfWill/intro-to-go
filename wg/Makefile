BIN_PATH = build/bin

.PHONY: all
all: rut

.PHONY: run
run:
	go run main.go

.PHONY: bench
bench:
	go test -bench=. -benchmem

.PHONY: test
test:
	go test -count=1 -cover ./...
