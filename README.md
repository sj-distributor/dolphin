# dolphin

### 目前处于开发阶段，本文档随时会更改，最新进度请切换develop分支了解

### [TODO代办事项](docs/TODO.md)

### 文件结构
```
  .
  ├── cmd  cli入口操作命令
  ├── docs  说明文档
  └── example  示例文件
      └── auth  权限验证相关
      └── gen  通过模板生成的代码目录
      └── src  自定义操作目录：扩展方法、处理功能逻辑
  ├── model  模型文件夹
  ├── templates  模板文件夹
  ├── makefile  自动化操作make配置文件
```

### 必做事项
```
1. 安装Golang，设置 GOPATH 环境变量
2. go get -d golang.org/x/tools/cmd/goimports
3. go install golang.org/x/tools/cmd/goimports
```

### 快速上手（执行过前面4步的，可从第5步开始）

1. `mkdir dolphin-test && cd dolphin-test`
2. `go mod init github.com/sj-distributor/dolphin-test`
3. `go get -d github.com/sj-distributor/dolphin`
4. `go run github.com/sj-distributor/dolphin init`
5. 修改 `model` 目录下的`graphql`的文件（可选）
6. 生成代码：`make generate`
7. 同步表：`make migrate`
8. 运行：`make start`

#### 可在 `makefile` 文件扩展 `make` 命令

### example 示例说明
  1. `cd example`
  2. 修改 `makefile` 文件的 `DATABASE` 值
  3. 修改`model`目录下的`graphql`的文件（可选）
  4. `make generate` 生成最新代码（`graphql`文件没改变可以不操作）
  5. `make migrate` 同步`graphql`数据表结构，没改变可以不用同步
  6. `make start` 启动项目

### model目录下的graphql文件说明
```
type Todo @entity(title: "代办事项") {
	title: String!
}
```
```
  .
  ├── Todo  这里指数据库表的命名
  └── @entity  指令
      └── title  Todo的描述说明
    └── title  表字段
```
## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2023 SJ Distributor
