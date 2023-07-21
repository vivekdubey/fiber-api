server:
	go run cmd/main.go

createmigration:
	migrate create -ext sql -dir internal/db/migration -seq $(migration_name)

migrateup:
	migrate -path internal/db/migration -database postgres://fiber_api_user:mypassword@localhost:5432/go_fiber_api -verbose up

migratedown:
	migrate -path internal/db/migration -database postgres://fiber_api_user:mypassword@localhost:5432/go_fiber_api -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: server migrateup migratedown test
