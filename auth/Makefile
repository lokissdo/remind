migrate_up:
	migrate -path db/migration -database "postgres://remind_user:haFaMrzhWE0sFRMgHqdRrd6V4lBZniLY@dpg-cpduaf7sc6pc73970qk0-a.oregon-postgres.render.com/remind" -verbose up $(level)

migrate_down:
	migrate -path db/migration -database "postgres://remind_user:haFaMrzhWE0sFRMgHqdRrd6V4lBZniLY@dpg-cpduaf7sc6pc73970qk0-a.oregon-postgres.render.com/remind" -verbose down $(level)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=remind_auth \
	proto/*.proto

redis:
	docker run --name redis -p 6379:6379 -d redis:7.2.5-alpine

server:
	go run main.go

.PHONY: migrate_up migrate_down new_migration sqlc test proto redis test server