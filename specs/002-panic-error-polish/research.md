# Research: Panic→Error 重构与代码完善

**Feature**: 002-panic-error-polish  
**Date**: 2026-03-03

## R1: panic 调用位置和调用链映射

**Decision**: 需要修改 6 处 panic，保留 1 处（`getNamedType`）

| 位置 | 当前行为 | 调用方 | 修改策略 |
|------|---------|--------|---------|
| `model/model.go:96` `Object()` | panic "not found" | `enrichment.go`, `model.object-field.go`, `model.object-relationship.go` | 返回 `(Object, error)` |
| `model/model.go:105` `ObjectExtension()` | panic "not found" | `model.object-field.go` | 返回 `(ObjectExtension, error)` |
| `model/model.object.go:157` `Relationship()` | panic "not found" | `model.object-field.go` | 返回 `(ObjectRelationship, error)` |
| `model/model.object-relationship.go:63` | panic "invalid value" | `model.object-relationship.go` 各方法 | 返回 `(string, error)` |
| `model/model.object-relationship.go:74` | panic "invalid value" | 同上 | 返回 `(bool, error)` |
| `model/model.object-relationship.go:81` | panic "missing inverse" | `InverseRelationship()` | 返回 `(ObjectRelationship, error)` |
| `model/config.go:57` | panic "parse error" | `cmd/gen.go` 模板数据 | 返回 `(time.Duration, error)` |

**保留**: `model/graphql-helpers.go:27` `getNamedType()` — 属于编程错误，传入无法解析类型意味着代码逻辑有 bug

**Rationale**: Go 惯例是只在真正不可恢复的编程错误中使用 panic。模型解析的"未找到"场景是用户输入错误，应返回 error

## R2: 签名变更的级联影响分析

**Decision**: 采用"从底向上"修改策略，先改最底层方法，再改调用方

调用链深度分析：
1. `StringForRelationshipDirectiveAttribute` → 被 `Target()`, `InverseRelationshipName()`, `JoinTable()` 等调用
2. `Object()` → 被 `enrichment.go`, `model.object-field.go`, 模板引用
3. 模板引用：`templates/*.go` 中的模板字符串**不受影响**（它们是生成到目标项目的代码，不调用 dolphin 的 model 包）

**Rationale**: 从底层开始修改可以保证中间状态始终可编译

## R3: 错误传播策略

**Decision**: 错误在 `model/` 层产生，通过 `enrichment.go` 传播到 `cmd/gen.go`，最终以 `cli.NewExitError` 输出

**Rationale**: 现有的 `EnrichModel()` 已返回 error，是天然的错误汇聚点。大部分 panic 发生在 enrichment 阶段或模板渲染阶段

## R4: 单元测试策略

**Decision**: 在 `model/` 和 `utils/` 包目录下直接创建 `_test.go` 文件（Go 惯例）

**Rationale**: Go 测试文件应与源码同目录。使用标准 `testing` 包，无需额外依赖。测试 fixture 使用内联 GraphQL 字符串

## R5: Go 版本和工具

**Version**: Go 1.24.0  
**Testing**: `go test ./...`  
**Linting**: `go vet ./...`
