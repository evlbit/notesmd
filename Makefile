build:
	@go build -o bin/notesmd cmd/api/main.go

run: build
	@./bin/notesmd