db:
	docker run --name=number-storage-db -e POSTGRES_PASSWORD='password12' -p 5432:5432 -d --pull=missing postgres

migrate:
	migrate -path ./migrations -database 'postgres://postgres:password12@localhost:5432/postgres?sslmode=disable' up

protobuf:
	protoc -I ./proto/ --go_out ./ --go-grpc_out ./ ./proto/number_storage.proto