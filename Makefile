lint:
	@golangci-lint run -c .golangci.yml --allow-parallel-runners

fmt:
	@find . -type f -name '*.go' -not -path "*.bm.go" -not -path "*.pb.go" -not -path "*wire_gen.go" | xargs goimports -w
