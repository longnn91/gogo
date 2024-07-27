mysql:
$ docker run --name gogo-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=be123@ -d mysql:tag

createdb:
$ docker exec -it gogo-mysql createdb --username=root --owner=root gogo-data

dropdb:
$ docker exec -it gogo-mysql dropdb gogo-data

migrateup:
$ migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3307)/gogo-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
$ migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3307)/gogo-data" -verbose down

.PHONY: mysql createdb dropdb migrateup migratedown