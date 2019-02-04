TARGET=comedian

all: fmt clean build docker

clean:
	rm -rf $(TARGET)

fmt:
	go fmt ./...

build:
	go build -o $(TARGET) main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o $(TARGET) main.go

build_docker:
	docker build -t comedian .

docker: build_linux build_docker

migration-ci:
            ./goose -dir migrations mysql "comedian:comedian@tcp(db:3306)/comedian"  up

migrate:
	goose -dir migrations mysql "comedian:comedian@tcp(db:3306)/comedian"  up

c_migration:
	goose -dir migrations create migration_name sql
	
db_clean:
	goose -dir migrations mysql "comedian:comedian@/comedian"  reset
	goose -dir migrations mysql "comedian:comedian@/comedian"  up

run_tests:
	go test ./bot/ ./notifier/ ./config/ ./api/ ./utils/ ./storage/ ./reporting/ ./collector ./model -cover

test: db_clean run_tests
