postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:alpine3.18

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root shoping_cart_db

dropdb:
	docker exec -it postgres-container dropdb shoping_cart_db


.PHONY: postgres, createdb, dropdb