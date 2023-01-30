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

### example 示例说明
  1. `cd example`
  2. 修改 `makefile` 文件的 `DATABASE` 值
  3. `make generate` 生成最新代码（`graphql`文件没改变可以不操作）
  4. `make migrate` 同步`graphql`数据表结构，没改变可以不用同步
  5. `make start` 启动项目

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2023 SJ Distributor
