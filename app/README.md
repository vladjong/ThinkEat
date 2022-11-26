docker run --name=think_eat -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres

migrate -path ./db/migrations -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up