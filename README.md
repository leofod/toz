## Запуск
```
docker-compose build && docker-compose up
```
## gRPC
```
docker exec -it {container_ID} /bin/sh
go run create_short_url.go
go run get_full_url.go
```