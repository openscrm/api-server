SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
cd ../
swag init
go build  -o bin/api-server main.go