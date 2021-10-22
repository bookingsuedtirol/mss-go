env:
	@grep -v '^#' examples/.env | xargs

simple:
	@go run examples/simple/*

advanced:
	@go run examples/advanced/*

test:
	go test -v
