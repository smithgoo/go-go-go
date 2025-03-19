go mod tidy

go mod vendor

go run main.go

数据库链接就行了 直接用ORM处理其他的逻辑不用处理 直接写方法 CURD


命令行测试 

post

    curl -X POST http://localhost:8080/v1/product/add -d '{"name": "Sample Product", "price": 100}' -H "Content-Type: application/json"

get

    curl http://localhost:8080/v1/product/0


前端校验 token的时候需要去掉引号

下面是测试
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwiZXhwIjoxNzI2Mjk0Mjg0fQ.VvKFR-1cV6AasyS84VYHZHa5w_l2gmCt7-nKhGqMLQg"  http://192.168.1.90:8080/user/home

