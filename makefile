db:
	docker run --name=number-storage-db -e POSTGRES_PASSWORD='password12' -p 5432:5432 -d --pull=missing postgres