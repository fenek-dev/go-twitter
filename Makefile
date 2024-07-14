PROJECTNAME=$(shell basename "$(PWD)")

local:
	make -j 5 db-up sso cache write-api read-api 

storage-up:
	make -j 2 db-up redis

protogen:
	export PATH="$PATH:$(go env GOPATH)/bin" && protoc -I proto proto/twitter.proto --go_out=./proto/protogen/ --go_opt=paths=source_relative --go-grpc_out=./proto/protogen/ --go-grpc_opt=paths=source_relative && mockgen -source=proto/protogen/twitter_grpc.pb.go -destination=proto/protogen/mocks/twitter_mock.go

# sso
sso:
	go run src/sso/cmd/main.go --config=src/sso/config/config_local.yaml

# write-api
write-api:
	go run src/write-api/cmd/main.go --config=src/write-api/config/config_local.yaml


# cache
cache:
	go run src/cache/cmd/main.go --config=src/cache/config/config_local.yaml

# read-api
read-api:
	go run src/read-api/cmd/main.go --config=src/read-api/config/config_local.yaml

# redis
redis-up:
	docker compose -f=docker/docker-compose.redis.yaml -p=redis_twitter up
redis-down:
	docker compose -f=docker/docker-compose.redis.yaml -p=redis_twitter down

# db
db-up:
	docker compose -f=docker/docker-compose.db.yaml -p=pg_twitter up
db-down:
	docker compose -f=docker/docker-compose.db.yaml -p=pg_twitter down
db-migrate:
	migrate -database "postgres://user:password@localhost:5432/twitter?sslmode=disable" -path migrations/pg up

jaeger:
	docker compose -f=docker/docker-compose.jaeger.yaml -p=jeager_twitter up

stress:
	oha -m PUT -H "Cookie: go_twitter_auth_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXNjcmlwdGlvbiI6IiIsImV4cCI6MTcyMDg4NTExOSwidXNlcm5hbWUiOiJEZXJlY2suUGZlZmZlciJ9.Li9LfzjBKAuje9b8ll_2xBnTlXMnjCFYiMvmZUN-Ajw" -H "Content-Type: application/json" -d '{"username":"Vernie88","content":"Similique sunt nesciunt."}' -c 200 -n 100000 http://localhost:8000/api/v1/tweet