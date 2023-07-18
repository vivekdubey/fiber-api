server:
	go run cmd/main.go

migrateup:
	migrate -path internal/db/migration -database postgres://fiber_api_user:mypassword@localhost:5432/go_fiber_api -verbose up

migratedown:
	migrate -path internal/db/migration -database postgres://fiber_api_user:mypassword@localhost:5432/go_fiber_api -verbose down

test:
	go test -v -cover ./...

.PHONY: server migrateup migratedown test
