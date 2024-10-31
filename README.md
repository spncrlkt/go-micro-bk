# gRPC microservices in Go


```
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=buttnutt -e MYSQL_DATABASE=order mysql

DATA_SOURCE_URL=root:buttnutt@tcp(127.0.0.1:3306)/order APPLICATION_PORT=3000 ENV=development go run cmd/main.go

grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' -plaintext localhost:3000 Order/Create

```

