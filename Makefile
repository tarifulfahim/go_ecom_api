build:
	@go build -o bin/ecom cmd/main.go


run: build
	@./bin/ecom