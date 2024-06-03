go mod tidy

go mod vendor

go run main.go

数据库链接就行了 直接用ORM处理其他的逻辑不用处理 直接写方法 CURD


命令行测试 

post

    curl -X POST http://localhost:8080/v1/product/add -d '{"name": "Sample Product", "price": 100}' -H "Content-Type: application/json"

get

    curl http://localhost:8080/v1/product/0
