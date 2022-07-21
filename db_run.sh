docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

## connect to postgres
# docker exec -it 30db055754df /bin/bash

## create migrate file
# migrate create -ext sql -dir ./migration -seq init

## apply migrate
# migrate -path ./migration -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up