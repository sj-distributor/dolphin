# DATABASE = sqlite3://test.db
DATABASE = mysql://'root:123456@localhost:3306/dolphin_example?charset=utf8mb4&parseTime=True&loc=Local'

PORT = 8080

mysql:
	docker run --restart=always -itd --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

redis:
	docker run --restart=always -itd --name redis -p 6379:6379 redis

docker:
	docker run --rm -it \
	--link mysql:db \
	--name golang \
	-e GOPATH=/data/go \
	-e GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
	-v ~/docker/go:/data/go \
	-v ~/wwwroot:/data/wwwroot \
	-p $(PORT):$(PORT) \
	golang bash

generate:
	GO111MODULE=on go run github.com/sj-distributor/dolphin

migrate:
	DATABASE_URL=$(DATABASE) PORT=$(PORT) go run main.go migrate

start:
	DATABASE_URL=$(DATABASE) PORT=$(PORT) DEBUG="true" go run ./main.go start --cors
