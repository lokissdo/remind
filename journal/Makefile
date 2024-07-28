proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=remind_journal \
	proto/*.proto
	
migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/journal?sslmode=disable" -verbose up $(level)

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/journal?sslmode=disable" -verbose down $(level)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: proto migrate_up migrate_down new_migration sqlc test server