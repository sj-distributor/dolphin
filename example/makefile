# DATABASE = sqlite3://test.db
DATABASE = mysql://'root:123456@localhost:3306/dolphin?charset=utf8mb4&parseTime=True&loc=Local'

PORT = 8080
DIR = go mod tidy

generate:
	GO111MODULE=on go run github.com/sj-distributor/dolphin

migrate:
	$(DIR) && DATABASE_URL=$(DATABASE) PORT=$(PORT) go run *.go migrate

start:
	$(DIR) && DATABASE_URL=$(DATABASE) PORT=$(PORT) DEBUG="true" go run ./server.go start --cors