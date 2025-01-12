.PHONY: test
test:
	go test -coverprofile=coverage.out ./...

.PHONY: cover
cover: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
