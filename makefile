db:
	docker run --name=number-storage-db -e POSTGRES_PASSWORD='password12' -p 5432:5432 -d --pull=missing postgres

migrate:
	migrate -path ./migrations -database 'postgres://postgres:password12@localhost:5432/postgres?sslmode=disable' up

protobuf:
	protoc -I ./proto --go_out ./ --go-grpc_out ./ --grpc-gateway_out ./ --openapiv2_out ./openapiv2 ./proto/number_storage.proto

swagger:
	docker run --name=swagger-ui -p 8008:8080 \
    -e SWAGGER_JSON=/openapiv2/number_storage.swagger.json \
    -v ${PWD}/openapiv2/:/openapiv2 \
    swaggerapi/swagger-ui