postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:alpine3.18

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root shoping_cart_db

dropdb:
	docker exec -it postgres-container dropdb shoping_cart_db

migrateup:
	 migrate -path ./db/migration -database "postgres://root:1234@localhost:5432/shoping_cart_db?sslmode=disable" -verbose up

migratedown:
	 migrate -path ./db/migration -database "postgres://root:1234@localhost:5432/shoping_cart_db?sslmode=disable" -verbose down



.PHONY: postgres, createdb, dropdb