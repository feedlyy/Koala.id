.PHONY: postgres migrate

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

migrate:
	migrate -source file://Rest/db/migrations \
 			-database postgres://root:secret@localhost/koala?sslmode=disable up

down:
	migrate -source file://Rest/db/migrations \
     			-database postgres://root:secret@localhost/koala?sslmode=disable down