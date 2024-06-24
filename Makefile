PROJECTNAME=$(shell basename "$(PWD)")

local:
	make -j 4 db-up sso write-api read-api

# sso
sso:
	go run src/sso/cmd/main.go --config=src/sso/config/config_local.yaml
sso-protogen:
	export PATH="$PATH:$(go env GOPATH)/bin" && protoc -I proto proto/*.proto --go_out=./src/sso/protogen/ --go_opt=paths=source_relative --go-grpc_out=./src/sso/protogen/ --go-grpc_opt=paths=source_relative

# write-api
write-api:
	go run src/write-api/cmd/main.go --config=src/write-api/config/config_local.yaml


# read-api
read-api:
	go run src/read-api/cmd/main.go --config=src/read-api/config/config_local.yaml

# redis
redis-up:
	docker-compose -f=storage/docker-compose.redis.yaml -p=redis_twitter up
redis-down:
	docker-compose -f=storage/docker-compose.redis.yaml -p=redis_twitter down

# db
db-up:
	docker-compose -f=storage/docker-compose.db.yaml -p=pg_twitter up
db-down:
	docker-compose -f=storage/docker-compose.db.yaml -p=pg_twitter down
db-migrate:
	migrate -database "postgres://user:password@localhost:5432/twitter?sslmode=disable" -path storage/migrations up