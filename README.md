# go-twitter

`docker-compose -f=db/docker-compose.yaml -p=pg_twitter up`

golang-migrate
`migrate -database "postgres://user:password@localhost:5432/twitter?sslmode=disable" -path db/migrations up`
