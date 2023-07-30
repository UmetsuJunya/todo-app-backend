# todo-app-backend

# How to use

## build from docker-compose.yml

docker-compose up --build

## make migrations

docker exec -it todo-backend-docker-go-1 sh
cd migration
go run main.go up

## Confirm DB table

docker exec -it todo-backend-docker-mysql-1 sh
mysql -u root -p
use backend;
show tables;
