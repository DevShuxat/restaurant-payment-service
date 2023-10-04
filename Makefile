PWD=$(shell pwd)
SERVICE=restaurant-svc
MIGRATION_PATH =${PWD}/src/infrastructure/migrations
PROTOS_PATH =./src/infrastructure/protos

server:
	go run main.go

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/restaurant?sslmode=disable&search_path=public" -path ./pkg/database/migrations up

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/restaurant?sslmode=disable&search_path=public" -path ./pkg/database/migrations down

add-protos-submodules:
	git submodule add git@github.com:DevShuxat/services-proto.git ./src/infrastructure/protos

pull-protos-submodules:
	git submodule update --recursive --remote


gen-restaurant-proto:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PROTOS_PATH)/restaurant \
	$(PROTOS_PATH)/restaurant/*.proto

docker: bin-lunix
	docker build --rm -t restaurant-svc -f ${PWD}./deploy/Dockerfile .

compose-up:
	docker-compose -f ./deploy/docker-compose.yml up

compose-down:
	docker-compose -f ./deploy/docker-compose.yml down