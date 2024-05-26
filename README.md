# go-twitter

`docker-compose -f=db/docker-compose.yaml -p=pg_twitter up`

golang-migrate
`migrate -database "postgres://user:password@localhost:5432/twitter?sslmode=disable" -path db/migrations up`

protoc
`protoc -I proto proto/*.proto --go_out=./src/auth/protogen/ --go_opt=paths=source_relative --go-grpc_out=./src/auth/protogen/ --go-grpc_opt=paths=source_relative`
