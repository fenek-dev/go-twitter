PROJECTNAME=$(shell basename "$(PWD)")

all:
	make -j 2 services db

services:
	make -j 3 auth tweet user

db:
	make -j 3 db-up redis-up cassandra-up

protogen:
	export PATH="$PATH:$(go env GOPATH)/bin" && protoc -I proto proto/twitter.proto --go_out=./proto/protogen/ --go_opt=paths=source_relative --go-grpc_out=./proto/protogen/ --go-grpc_opt=paths=source_relative

mockgen:
	mockery --dir=proto/protogen --output=proto/protogen/mocks --outpkg=proto_mocks --all && \
	cd src && \
	mockery --dir=auth --output=auth/mocks --outpkg=mocks --all && \
	mockery --dir=tweet --output=tweet/mocks --outpkg=mocks --all && \
	mockery --dir=user --output=user/mocks --outpkg=mocks --all

auth:
	cd ./src/auth && air -- --config=config/config_local.yaml

tweet:
	cd ./src/tweet && air -- --config=config/config_local.yaml

user:
	cd ./src/user && air -- --config=config/config_local.yaml

cassandra-up:
	docker compose -f=docker/docker-compose.cassandra.yaml -p=cassandra_twitter up
cassandra-down:
	docker compose -f=docker/docker-compose.cassandra.yaml -p=cassandra_twitter down
cassandra-migrate:
	docker cp migrations/cassandra/init.cql cassandra1:/init.cql && \
	docker exec -it cassandra1 cqlsh -f /init.cql

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