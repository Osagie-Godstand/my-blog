build-app:
	@go build -o bin/app ./cmd/app/

run: build-app
	@./bin/app

docker:
	@echo "building Docker container"
	@docker build -t app -f Dockerfile .
	@echo "running API inside Docker container"
	@docker run -p 8080:8080 app
