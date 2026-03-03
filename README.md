# dolphin

GraphQL 代码生成器 —— 从 `.graphql` 模型文件自动生成 Go 后端代码（GORM 模型、GraphQL resolver、CRUD 接口、校验器等）。

## 文件结构

```
├── cmd/           CLI 入口命令
├── model/         模型解析与代码生成逻辑
├── templates/     Go 模板文件（生成目标代码）
├── example/       示例项目
│   ├── auth/      权限验证
│   ├── gen/       生成的代码目录
│   └── src/       自定义业务逻辑
├── gqlgen/        gqlgen 依赖
└── utils/         工具函数
```

## 快速上手

### 前置条件

```bash
# 1. 安装 Go，设置 GOPATH
# 2. 安装 goimports
go install golang.org/x/tools/cmd/goimports@latest
```

### 创建新项目

```bash
mkdir myproject && cd myproject
go mod init github.com/yourname/myproject
printf '//go:build tools\npackage tools\nimport (_ "github.com/sj-distributor/dolphin"\n _ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
go run github.com/sj-distributor/dolphin init
```

### 日常开发

```bash
make generate   # 根据 model/*.graphql 生成代码
make migrate    # 同步数据库表结构
make start      # 启动服务
```

## GraphQL 模型语法

### 实体定义

```graphql
type User @entity(title: "用户管理") {
  phone: String! @column(gorm: "type:varchar(32);NOT NULL;index:phone;") @validator(required: "true", type: "phone")
  password: String! @column(gorm: "type:varchar(64);NOT NULL;") @validator(required: "true", type: "password")
  nickname: String @column(gorm: "type:varchar(64);DEFAULT NULL;")
  age: Int @column(gorm: "type:int(3);default:1;") @validator(type: "int")
}
```

| 指令 | 作用 | 位置 |
|------|------|------|
| `@entity(title: "...")` | 声明数据库实体，title 为表描述 | type 上 |
| `@column(gorm: "...")` | 数据库列定义（GORM 标签语法） | 字段上 |
| `@validator(...)` | 字段校验规则 | 字段上 |
| `@hasRole(role: ADMIN)` | 访问权限控制 | type / 字段上 |

### 系统自动生成的字段

每个 `@entity` 自动包含以下字段，无需手动定义：

| 字段 | 类型 | 说明 |
|------|------|------|
| `id` | varchar(36) | UUID 主键 |
| `createdAt` | bigint(13) | 创建时间（毫秒） |
| `updatedAt` | bigint(13) | 更新时间（毫秒） |
| `deletedAt` | bigint(13) | 删除时间（软删除） |
| `createdBy` | varchar(36) | 创建人 |
| `updatedBy` | varchar(36) | 更新人 |
| `deletedBy` | varchar(36) | 删除人 |
| `isDelete` | int(2) | 是否删除：1/正常、2/删除 |
| `weight` | int(2) | 权重（排序用） |
| `state` | int(2) | 状态：1/正常、2/禁用 |

## @validator 校验指令

### 参数列表

| 参数 | 类型 | 说明 | 示例 |
|------|------|------|------|
| `required` | String | 是否必填 | `required: "true"` |
| `immutable` | String | 创建后不可修改 | `immutable: "true"` |
| `type` | String | 正则校验类型（见 `utils/rule.go`） | `type: "phone"` |
| `minLength` | Int | 最小字符串长度 | `minLength: 6` |
| `maxLength` | Int | 最大字符串长度 | `maxLength: 32` |
| `minValue` | Int | 最小数值 | `minValue: 0` |
| `maxValue` | Int | 最大数值 | `maxValue: 100` |
| `unique` | String | 是否唯一（不允许重复） | `unique: "true"` |
| `uniqueScope` | String | 唯一性条件范围字段（可选） | `uniqueScope: "uid"` |

### 内置校验类型 (`type`)

在 `utils/rule.go` 中定义，常用类型：

| 类型 | 说明 |
|------|------|
| `phone` | 手机号格式 |
| `email` | 邮箱格式 |
| `password` | 密码格式（自动加密） |
| `int` | 整数格式 |
| `justInt` | 纯数字格式 |

### 唯一性校验

#### 简单唯一

字段值在整个表中不允许重复：

```graphql
type User @entity(title: "用户管理") {
  phone: String! @column(gorm: "type:varchar(32);NOT NULL;") @validator(required: "true", unique: "true")
}
```

#### 条件唯一（带范围）

字段值在指定范围内不允许重复（例如同一 `uid` 下 `name` 不能重复）：

```graphql
type Product @entity(title: "产品管理") {
  uid: ID @column(gorm: "type:varchar(36);")
  name: String! @column(gorm: "type:varchar(64);NOT NULL;") @validator(required: "true", unique: "true", uniqueScope: "uid")
}
```

上述定义表示：
- `name` 字段值在**同一 `uid`** 下不允许重复
- 不同 `uid` 下可以有相同的 `name`
- `uniqueScope` 为可选参数，省略时按全表唯一

> **注意**：唯一性校验依赖 context 中注入的 `db`（*gorm.DB）和 `tableName`（string）。校验时自动排除软删除记录（`is_delete = 1`）。

## 生成的代码结构

运行 `make generate` 后，会在以下目录生成代码：

```
├── gen/                        自动生成（勿手动修改）
│   ├── schema.graphqls         完整 GraphQL Schema
│   ├── generated.go            gqlgen 生成的 resolver
│   ├── models.go               Go 结构体（含 GORM 标签）
│   ├── database.go             数据库初始化
│   ├── resolver-mutations.go   Mutation resolver
│   ├── resolver-queries.go     Query resolver
│   └── ...
├── src/                        用户自定义（可修改）
│   ├── resolver.go             自定义 resolver 入口
│   ├── resolver_gen.go         生成的 resolver 骨架
│   └── extend.go               自定义 Query/Mutation
├── utils/                      自动生成（勿手动修改）
│   ├── validator.go            校验器（由模板生成）
│   ├── rule.go                 校验规则定义
│   ├── utils.go                工具函数
│   └── encrypt.go              加密工具
```

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2023 SJ Distributor
