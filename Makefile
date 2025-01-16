mysql:
	docker run --name go-restaurant-mysql -p 3308:3306 -e MYSQL_ROOT_PASSWORD=be123@ -d mysql:latest

createdb:
	docker exec -it go-restaurant-mysql mysql -uroot -pbe123@ -e "CREATE DATABASE IF NOT EXISTS \`go-restaurant-data\`"

dropdb:
	docker exec -it go-restaurant-mysql mysql -uroot -pbe123@ -e "DROP DATABASE IF EXISTS \`go-restaurant-data\`"

migrateup:
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data" -verbose down

.PHONY: mysql createdb dropdb migrateup migratedown
