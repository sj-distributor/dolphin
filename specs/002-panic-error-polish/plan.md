# Implementation Plan: Panic→Error 重构与代码完善

**Branch**: `002-panic-error-polish` | **Date**: 2026-03-03 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/002-panic-error-polish/spec.md`

## Summary

将 dolphin 代码生成器核心 `model/` 包中的 6 处 panic 替换为 Go 惯例的 error 返回模式，确保错误沿调用链传播至 CLI 层。同时为所有导出符号添加 godoc 规范注释，并为核心函数添加单元测试。

## Technical Context

**Language/Version**: Go 1.24.0  
**Primary Dependencies**: gqlgen v0.17.85, urfave/cli v1.22.15  
**Storage**: N/A  
**Testing**: Go 标准 `testing` 包 + `go test`  
**Target Platform**: CLI 工具  
**Project Type**: single  
**Constraints**: 正常场景行为等价（生成代码不变）

## Constitution Check

| Gate | Status | Notes |
|------|--------|-------|
| SDD 协议 | ✅ PASS | spec → plan → tasks 流程 |
| Monorepo Context | ✅ PASS | 改动对象是 dolphin 工具本身 |
| Dolphin Engine Rule | ✅ PASS | 不引入新模式 |
| Code Integrity | ✅ PASS | 消除 panic 正是该原则的体现 |

## Project Structure

### Source Code (变更范围)

```text
model/
├── model.go                     # Object(), ObjectExtension() 签名变更
├── model.object.go              # Relationship() 签名变更
├── model.object-field.go        # 更新调用方
├── model.object-relationship.go # 3 处 panic → error
├── config.go                    # ConnMaxLifetime() 签名变更
├── enrichment.go                # 更新调用方（错误汇聚点）
├── utils.go                     # RegexpReplace 错误处理
├── graphql-helpers.go           # 保留 panic（编程错误）
├── model_test.go                # [NEW] 核心方法测试
├── utils_test.go                # [NEW] 工具函数测试
└── object_field_test.go         # [NEW] 字段方法测试

templates/
└── helpers.go                   # prompt() panic → error

utils/
└── file_utils_test.go           # [NEW] 文件工具测试

cmd/
├── root.go                      # 注释规范化
└── gen.go                       # 注释规范化
```

---

## 变更清单

### Phase 1: 底层方法签名变更 (US1 — P1 MVP)

#### [MODIFY] [model.object-relationship.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.object-relationship.go)

**修改 3 处 panic → error（最底层，无外部调用方依赖）**

1. `StringForRelationshipDirectiveAttribute(name)`: `string` → `(string, error)`
   - 第63行 panic → `return "", fmt.Errorf(...)`
2. `BoolForRelationshipDirectiveAttribute(name)`: `bool` → `(bool, error)`
   - 第74行 panic → `return false, fmt.Errorf(...)`
3. `InverseRelationshipName()`: `string` → `(string, error)`
   - 第81行 panic → `return "", fmt.Errorf(...)`

**更新内部调用方（同文件）**：
- `Target()` → 处理 `StringForRelationshipDirectiveAttribute` 的 error
- `InverseRelationship()` → 处理 `InverseRelationshipName` 的 error
- `JoinTable()` → 处理 error
- `IsManyToMany()`, `IsOneToOne()`, `IsManyToOne()` 等 → 处理 `BoolForRelationshipDirectiveAttribute` 的 error

#### [MODIFY] [model.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.go)

1. `Object(name)`: `Object` → `(Object, error)`
   - 第96行 panic → `return Object{}, fmt.Errorf(...)`
2. `ObjectExtension(name)`: `ObjectExtension` → `(ObjectExtension, error)`
   - 第105行 panic → `return ObjectExtension{}, fmt.Errorf(...)`

#### [MODIFY] [model.object.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.object.go)

1. `Relationship(name)`: `ObjectRelationship` → `(ObjectRelationship, error)`
   - 第157行 panic → `return ObjectRelationship{}, fmt.Errorf(...)`

#### [MODIFY] [config.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/config.go)

1. `ConnMaxLifetime()`: `time.Duration` → `(time.Duration, error)`
   - 第57行 panic → `return 0, fmt.Errorf(...)`

#### [MODIFY] [utils.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/utils.go)

1. `RegexpReplace(src, expr, repl)`: 添加 `regexp.Compile` 错误处理
   - 编译失败时返回原字符串（保持兼容性）

#### [MODIFY] [helpers.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/helpers.go)

1. `prompt()` 中密码读取 panic → 返回 error

### Phase 2: 调用方更新 (US1)

#### [MODIFY] [model.object-field.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.object-field.go)

更新所有调用 `Object()`、`ObjectExtension()`、`Relationship()` 的方法以处理 error 返回。涉及方法：
- `TargetObject()`, `TargetObjectExtension()`, `HasTargetObject()`
- `HasTargetObjectExtension()`, `HasTargetTypeWithIDField()`

#### [MODIFY] [enrichment.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/enrichment.go)

更新 `EnrichModelObjects()` 中调用 `Object()` 和 relationship 方法的地方以处理 error。这是错误从 model 层传播到 cmd 层的关键汇聚点。

#### [MODIFY] [gen.go](file:///Users/marlon.m/wwwroot/triton/dolphin/cmd/gen.go)

确认 `EnrichModel()` 的 error 已正确传播到 `cli.NewExitError()`。

### Phase 3: 注释规范化 (US2)

#### [MODIFY] 多个文件

为以下文件中所有导出符号添加 godoc 规范注释：
- `cmd/root.go`: `Execute`
- `model/parser.go`: `Parse`
- `model/enrichment.go`: `EnrichModelObjects`, `EnrichModel`
- `model/model.go`: `SecretKey`, `Objects`, `HasObject`, `Object`, `ObjectExtension`, `ObjectEntities`
- `model/model.object.go`: 所有导出方法
- `model/graphql-helpers.go`: 所有函数
- `templates/helpers.go`: `WriteTemplate`, `WriteTemplateRaw`, `WriterOriginalFile`, `WriteInterfaceTemplate`

### Phase 4: 单元测试 (US3)

#### [NEW] [model/utils_test.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/utils_test.go)

测试用例：
- `TestGetRandomString` — 长度正确、字符在合法范围内
- `TestRegexpReplace` — 正常替换、无效正则返回原字符串
- `TestIndexOf` — 找到/未找到

#### [NEW] [utils/file_utils_test.go](file:///Users/marlon.m/wwwroot/triton/dolphin/utils/file_utils_test.go)

测试用例：
- `TestEnsureDir` — 创建新目录、已存在目录
- `TestFileExists` — 存在的文件、不存在的文件、目录

#### [NEW] [model/model_test.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model_test.go)

测试用例：
- `TestParse` — 基本 GraphQL schema 解析
- `TestObject` — 找到对象、未找到返回 error
- `TestObjectEntities` — 过滤非实体对象

---

## Verification Plan

### Automated Tests

**命令 1**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go build ./...`
- 验证所有签名变更和调用方更新编译通过

**命令 2**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go vet ./...`
- 验证无 vet 警告

**命令 3**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go test ./model/ ./utils/`
- 运行所有新增单元测试

**命令 4**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go test -cover ./model/`
- 验证 model 包覆盖率 ≥ 40%

### 行为等价性验证

在 `example/` 目录运行 `go run . generate`，确认正常模型的生成代码无变化。

### Manual Verification

构造一个引用不存在类型的 `.graphql` 文件，运行 `dolphin generate`，确认输出描述性错误信息而非 panic 崩溃。

## Complexity Tracking

> 无 Constitution 违规。
