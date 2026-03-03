# Implementation Plan: 代码质量重构

**Branch**: `001-code-quality-refactor` | **Date**: 2026-03-03 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-code-quality-refactor/spec.md`

## Summary

对 dolphin 代码生成器进行全面的代码质量重构。重点包括：清除死代码、修复已知 Bug、消除重复代码、升级废弃 API、改进错误处理策略、规范化代码注释、优化模板结构、添加单元测试基础设施。所有变更必须保证行为等价性——重构后生成的代码与重构前完全一致。

## Technical Context

**Language/Version**: Go 1.24.0  
**Primary Dependencies**: gqlgen v0.17.85, urfave/cli v1.22.15, graphql-go/graphql, iancoleman/strcase, jinzhu/inflection  
**Storage**: N/A（代码生成器，不涉及数据库）  
**Testing**: Go 标准 `testing` 包 + `go test`（当前无任何测试）  
**Target Platform**: CLI 工具，跨平台  
**Project Type**: single（CLI 代码生成器）  
**Performance Goals**: N/A  
**Constraints**: 重构后行为等价（生成的代码必须与重构前一致）  
**Scale/Scope**: ~70 个源文件，代码生成器核心

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Gate | Status | Notes |
|------|--------|-------|
| SDD 协议 | ✅ PASS | spec.md → plan.md → tasks.md 流程 |
| Monorepo Context | ✅ PASS | 本次重构针对 dolphin 工具本身，非 `/engine` 或 `/web` |
| Dolphin Engine Rule | ✅ PASS | 未引入新模式，仅优化现有代码 |
| Version Sync | ✅ PASS | Go 1.24.0，不升级依赖版本 |
| Code Integrity | ✅ PASS | 修复 Bug、提升代码质量正是该原则的体现 |

## Project Structure

### Documentation (this feature)

```text
specs/001-code-quality-refactor/
├── plan.md              # 本文件
├── spec.md              # 特性规范
├── research.md          # Phase 0 研究输出
├── data-model.md        # Phase 1 数据模型（本特性无实体，仅记录变更映射）
├── quickstart.md        # Phase 1 快速上手
├── checklists/
│   └── requirements.md  # 规范质量检查清单
└── tasks.md             # Phase 2 输出（由 /speckit.tasks 生成）
```

### Source Code (repository root)

```text
# Dolphin 代码生成器（重构范围）
cmd/
├── root.go              # CLI 入口 — 错误处理改进
├── gen.go               # 代码生成 — 死代码清除、config 加载优化
└── init.go              # 项目初始化 — 死代码清除、config 加载优化

model/
├── model.go             # 核心模型 — panic→error
├── model.object.go      # 对象定义 — panic→error
├── model.object-field.go        # 字段定义 — 注释规范化
├── model.object-field-input.go  # 输入字段 — Required() Bug 修复
├── model.object-relationship.go # 关系定义 — panic→error
├── config.go            # 配置 — panic→error
├── enrichment.go        # 模型丰富 — 注释规范化
├── printer.go           # Schema 打印 — 错误处理
├── parser.go            # 解析器 — 注释补全
├── definition.*.go      # 各定义文件 — 死代码清除
├── generics.go          # 泛型工具 — 无变更
├── graphql-helpers.go   # GQL 辅助 — 无变更
├── enums.go             # 枚举常量 — 无变更
└── utils.go             # 工具函数 — 错误处理、去重

templates/
├── helpers.go           # 模板辅助 — ioutil 升级、isExist Bug 修复、去重
├── resolver-mutations.go # Mutation 模板 — 注释规范化
├── resolver-queries.go  # Query 模板 — 注释规范化
├── resolver-utils.go    # 工具模板 — 无逻辑变更（模板代码）
├── utils-validator.go   # 验证器模板 — Max/Min Bug 修复
└── [其他 37 个模板文件]  # 注释规范化

utils/
└── file_utils.go        # 文件工具 — 无变更

tools/
└── run.go               # 命令运行 — 去重，合并到 templates 或 utils

tests/                   # 新增目录
├── model/
│   ├── model_test.go
│   ├── utils_test.go
│   └── object_field_test.go
└── utils/
    └── file_utils_test.go
```

**Structure Decision**: 保持现有目录结构，仅在根目录新增 `tests/` 目录存放单元测试（或直接在对应包目录添加 `_test.go` 文件，按 Go 惯例）。

---

## 变更清单

### Phase 1: 修复代码缺陷 (P1)

#### [MODIFY] [helpers.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/helpers.go)
- **L176**: 修复 `isExist()` 函数：`os.IsExist(err)` → `!os.IsNotExist(err)`
- **L8**: 移除 `io/ioutil` import，使用 `os.WriteFile`/`os.ReadFile`
- **L60,69,95**: `ioutil.WriteFile` → `os.WriteFile`
- **L60,69,95**: 文件权限 `0777` → `0644`
- **L42-52**: 移除注释掉的示例代码块
- **L116**: 移除注释掉的 `set -o pipefail` 行

#### [MODIFY] [model.object-field-input.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.object-field-input.go)
- **L32-38**: 简化 `Required()` 方法为 `return isNonNullType(o.Def.Type)`

#### [MODIFY] [utils-validator.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/utils-validator.go)
- **L160**: 修复 `Max()` 调用参数：`Max(intValue, *minValue)` → `Max(intValue, *maxValue)`
- **L124**: 修复 `Min()` 格式化：`"must be at least %s" + convertor.ToString(min)` → `fmt.Sprintf("must be at least %s", convertor.ToString(min))`
- **L137**: 修复 `Max()` 格式化：同上修复模式

#### [MODIFY] [printer.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/printer.go)
- **L107**: 添加 `printer.Print` 类型断言错误处理

### Phase 2: 死代码清除 (P1)

#### [MODIFY] [gen.go](file:///Users/marlon.m/wwwroot/triton/dolphin/cmd/gen.go)
- **L123-126**: 移除注释掉的 `model.BuildFederatedModel` 调用
- **L233-235**: 移除注释掉的 `ResolverSrcContext` 写入

#### [MODIFY] [init.go](file:///Users/marlon.m/wwwroot/triton/dolphin/cmd/init.go)
- **L72-74**: 移除注释掉的 `createDockerFile` 调用
- **L144-146**: 移除注释掉的 `UploadModel` 写入
- **L162-169**: 移除注释掉的 `createDockerFile` 函数
- **L233-235**: 移除注释掉的 `UpLoad` 写入

#### [MODIFY] [definition.inputs.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/definition.inputs.go)
- **L85-103**: 移除注释掉的旧条件逻辑代码块

#### [MODIFY] [definition.filter.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/definition.filter.go)
- **L18**: 移除注释掉的 `fields = append(fields, filterInputValues("id",...))` 行
- **L52-55**: 移除注释掉的 Description 定义

#### [MODIFY] [definition.mutation.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/definition.mutation.go)
- **L38, L62, L95, L100, L112**: 移除注释掉的 Description 字段

#### [MODIFY] [definition.query.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/definition.query.go)
- **L12, L46, L53, L60, L100**: 移除注释掉的代码

#### [MODIFY] [resolver-utils.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/resolver-utils.go) (模板)
- **L74-77**: 移除注释掉的 EventController 代码块

#### [MODIFY] [resolver-queries.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/resolver-queries.go) (模板)
- **L213-217**: 移除注释掉的 select/where 查询代码块

#### [MODIFY] [enrichment.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/enrichment.go)
- **L26**: 移除注释掉的 `columnDefinition` 行
- **L65**: 移除注释掉的 `createFederationServiceObject` 调用

### Phase 3: 消除重复代码与 Config 优化 (P2)

#### [MODIFY] [init.go](file:///Users/marlon.m/wwwroot/triton/dolphin/cmd/init.go)
- 重构 `initCmd.Action` 函数：将 `model.LoadConfigFromPath(p)` 改为在顶层加载一次，传递给所有 `create*File` 函数
- 调整所有 `create*File(p string)` 签名为 `create*File(p string, c *model.Config)`
- 移除每个函数内部重复的 `model.LoadConfigFromPath` 调用

#### [DELETE] [run.go](file:///Users/marlon.m/wwwroot/triton/dolphin/tools/run.go)
- 将 `RunInteractiveInDir` 统一到 `templates/helpers.go`（已有同名函数）
- 同时将 `tools` 包的唯一引用 `cmd/gen.go:159` 改为使用 `templates.RunInteractiveInDir`

#### [MODIFY] [gen.go](file:///Users/marlon.m/wwwroot/triton/dolphin/cmd/gen.go)
- **L159**: `tools.RunInteractiveInDir(...)` → `templates.RunInteractiveInDir(...)`
- 移除 `tools` import

### Phase 4: 错误处理改进 (P2)

#### [MODIFY] [model.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.go)
- **L91-97**: `Object(name)` 从 `panic` 改为返回 `(Object, error)`
- **L99-106**: `ObjectExtension(name)` 从 `panic` 改为返回 `(ObjectExtension, error)`
- 更新所有调用方以处理新的 error 返回

#### [MODIFY] [model.object-relationship.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/model.object-relationship.go)
- **L63, L74**: `StringForRelationshipDirectiveAttribute` / `BoolForRelationshipDirectiveAttribute` 中 `panic` → `error` 返回
- **L81**: `InverseRelationshipName` 中 `panic` → `error` 返回
- 更新所有调用方

#### [MODIFY] [config.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/config.go)
- **L57**: `ConnMaxLifetime()` 中 `panic` → 返回 `error`

#### [MODIFY] [utils.go](file:///Users/marlon.m/wwwroot/triton/dolphin/model/utils.go)
- **L34**: `RegexpReplace` 添加 `regexp.Compile` 错误处理

#### [MODIFY] [helpers.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/helpers.go)
- **L151**: `prompt` 函数中密码读取 `panic` → 返回 `error`

### Phase 5: 注释规范化 (P3)

#### [MODIFY] 多个文件
- 将 `// Execute ...`、`// Parse`、`// EnrichModelObjects ...` 等不完整注释替换为 godoc 规范注释
- 确保所有导出函数有描述性注释
- 统一注释语言风格（中文保持一致）
- 涉及文件：`cmd/root.go`、`model/parser.go`、`model/enrichment.go`、`model/model.go`、`model/model.object.go`、`model/graphql-helpers.go`、`tools/run.go`、`templates/helpers.go`

### Phase 6: 单元测试 (P3)

#### [NEW] model/model_test.go
- 测试 `Parse()`、`Objects()`、`HasObject()`、`ObjectEntities()` 等核心方法

#### [NEW] model/utils_test.go  
- 测试 `GetRandomString()`、`RegexpReplace()`、`IndexOf()`

#### [NEW] model/object_field_test.go
- 测试 `GoTypeWithPointer()`、`IsCreatable()`、`IsUpdatable()`、`FilterMapping()` 等关键方法

#### [NEW] utils/file_utils_test.go
- 测试 `EnsureDir()`、`FileExists()`

---

## Verification Plan

### Automated Tests

**命令**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go test ./...`
- 运行所有新增的单元测试
- 验证无编译错误和运行时失败

**命令**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go vet ./...`
- 验证无 vet 警告（包括废弃 API 检查）

**命令**: `cd /Users/marlon.m/wwwroot/triton/dolphin && go build ./...`
- 验证重构后项目可以成功编译

### 行为等价性验证

**步骤**:
1. 在重构前（`master` 分支），对 `example/` 项目运行 `go run . generate`，将 `example/gen/` 和 `example/utils/` 的生成代码备份到 `/tmp/dolphin-before/`
2. 切回 `001-code-quality-refactor` 分支，对同一 `example/` 运行 `go run . generate`
3. 使用 `diff -r /tmp/dolphin-before/ example/gen/` 比对生成代码
4. 差异应为零（模板的 Bug 修复除外，Bug 修复处需标注为预期变更）

### Manual Verification

请在重构完成后执行以下步骤：
1. 在现有业务项目中运行重构后的 dolphin `generate` 命令
2. 确认生成的代码无编译错误
3. 确认生成的 API 服务可正常启动

## Complexity Tracking

> 无 Constitution 违规，无需记录。
