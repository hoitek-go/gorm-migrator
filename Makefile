testcov:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
test:
	go test ./... -v
container-up:
	docker-compose -f ./docker-compose.test.yml up -d
container-down:
	docker-compose -f ./docker-compose.test.yml down