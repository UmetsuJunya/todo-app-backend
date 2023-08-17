# todo-app-backend

# How to use

## build from docker-compose.yml

docker-compose up --build

## make migrations

docker exec -it goDockerAPI sh<br>
cd migration<br>
go run main.go up

## Confirm DB table

docker exec -it goDockerDB sh<br>
mysql -u root -p<br>
show databases;<br>
use backend;<br>
show tables;
