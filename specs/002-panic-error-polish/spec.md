# Feature Specification: Panic→Error 重构与代码完善

**Feature Branch**: `002-panic-error-polish`  
**Created**: 2026-03-03  
**Status**: Draft  
**Input**: US5: panic → error 返回（影响 20+ 调用方）; US6: 注释规范化; US8: 单元测试

## User Scenarios & Testing

### User Story 1 - 错误处理改进：panic → error 返回 (Priority: P1) 🎯 MVP

作为使用 dolphin 代码生成器的开发者，我希望当输入的 GraphQL 模型存在错误（如引用不存在的类型、缺失关系定义）时，工具输出清晰的错误信息而不是直接崩溃（panic），以便我能快速定位并修复模型定义问题。

**Why this priority**: panic 导致程序崩溃且无堆栈上下文提示，用户体验极差且无法优雅恢复，是影响工具可用性的首要问题。

**Independent Test**: 构造一个引用不存在类型的 `.graphql` 文件，运行 `dolphin generate`，预期输出错误信息（非崩溃）。

**Acceptance Scenarios**:

1. **Given** 模型中引用了不存在的对象名, **When** 运行 generate 命令, **Then** 输出 "object 'XXX' not found in model" 错误信息并正常退出（exit code 1）
2. **Given** 关系定义缺少 inverse 属性, **When** 运行 generate 命令, **Then** 输出 "missing inverse value for X→Y relationship" 错误信息
3. **Given** 关系指令属性值无效, **When** 运行 generate 命令, **Then** 输出 "invalid XXX value for X→Y relationship" 错误信息
4. **Given** dolphin.yml 中 connMaxLifetime 格式无效, **When** 运行任何命令, **Then** 输出 "failed to parse config connMaxLifetime" 错误信息
5. **Given** 正常的模型文件, **When** 运行 generate 命令, **Then** 行为与重构前完全一致

---

### User Story 2 - 代码注释规范化 (Priority: P2)

作为 dolphin 项目的贡献者，我希望所有导出函数都有符合 Go 惯例（godoc）的文档注释，以便在 IDE 中悬停查看时能理解每个函数的用途。

**Why this priority**: 高质量注释提升开发体验和代码可维护性，但不影响功能正确性。

**Independent Test**: 运行 Go linter 确认所有导出符号有文档注释，无 "exported function XXX should have comment" 警告。

**Acceptance Scenarios**:

1. **Given** `cmd/root.go` 中 `Execute` 函数注释为 `// Execute ...`, **When** 检查注释, **Then** 注释描述了函数的实际用途
2. **Given** `model/parser.go` 中 `Parse` 函数注释不完整, **When** 检查注释, **Then** 注释描述了解析器输入、输出和功能
3. **Given** `model/enrichment.go` 中函数注释为 `// EnrichModelObjects ...`, **When** 检查注释, **Then** 注释说明了丰富逻辑的具体行为
4. **Given** 所有导出函数, **When** 运行 linter, **Then** 无 "should have comment" 类警告

---

### User Story 3 - 单元测试基础设施 (Priority: P3)

作为 dolphin 项目的维护者，我希望核心解析和工具函数有单元测试覆盖，以便在未来进行代码变更时有安全网。

**Why this priority**: 测试是长期可维护性的基础，但不改变产品功能。

**Independent Test**: 运行 `go test ./...` 全部通过，`model/` 包覆盖率 ≥ 40%。

**Acceptance Scenarios**:

1. **Given** `model/utils.go` 中的 `GetRandomString`、`RegexpReplace`、`IndexOf`, **When** 运行测试, **Then** 边界条件和常规用例均通过
2. **Given** `utils/file_utils.go` 中的 `EnsureDir`、`FileExists`, **When** 运行测试, **Then** 文件存在/不存在/目录创建场景均通过
3. **Given** `model/model.go` 中的 `Parse`、`Objects`、`HasObject`, **When** 运行测试, **Then** 基本 GraphQL 解析和对象查找功能通过
4. **Given** 重构后的 error 返回方法, **When** 传入无效参数, **Then** 返回预期的 error 而非 panic

---

### Edge Cases

- 修改 `Object()` 签名后，模板代码（`templates/`）中调用该方法的模板字符串是否需要同步更新？
- `getNamedType()` 中的 panic 是否属于真正不可恢复的编程错误（保留 panic 的合理场景）？
- 错误消息中是否需要包含源文件行号信息以方便用户定位？
- 同时修改多个方法签名时，中间状态能否保持编译通过？

## Requirements

### Functional Requirements

- **FR-001**: `model.Object(name)` 方法在对象不存在时必须返回 error 而非 panic
- **FR-002**: `model.ObjectExtension(name)` 方法在扩展不存在时必须返回 error 而非 panic
- **FR-003**: `ObjectRelationship` 的 `StringForRelationshipDirectiveAttribute()` 和 `BoolForRelationshipDirectiveAttribute()` 在属性无效时必须返回 error 而非 panic
- **FR-004**: `InverseRelationshipName()` 在缺少 inverse 定义时必须返回 error 而非 panic
- **FR-005**: `Config.ConnMaxLifetime()` 在解析失败时必须返回 error 而非 panic
- **FR-006**: `Object.Relationship(name)` 在关系不存在时必须返回 error 而非 panic
- **FR-007**: 所有 error 必须沿调用链正确传播至 CLI 层（`cmd/gen.go`），CLI 层输出错误信息并以 exit code 1 退出
- **FR-008**: 所有导出函数必须有符合 godoc 规范的描述性注释
- **FR-009**: `model/` 包核心函数必须有单元测试覆盖
- **FR-010**: `utils/` 包工具函数必须有单元测试覆盖
- **FR-011**: 正常场景下（模型无错误），行为与重构前完全一致

### Assumptions

- `graphql-helpers.go` 中 `getNamedType()` 的 panic 属于编程错误（不应传入无法解析的类型），保留 panic
- `templates/helpers.go` 中 `prompt()` 的密码读取 panic 改为返回 error
- `model/utils.go` 中 `RegexpReplace` 的 `regexp.Compile` 错误应返回原字符串而非 panic（保持兼容性）
- 单元测试使用 Go 标准 `testing` 包，无需引入额外测试框架

## Success Criteria

### Measurable Outcomes

- **SC-001**: dolphin 核心代码（排除 `example/`、`gqlgen/`）中 panic 调用从当前 7 处降至 ≤ 1 处（仅保留 `getNamedType` 的编程错误 panic）
- **SC-002**: 所有错误场景输出描述性错误信息且程序正常退出（exit code 1）
- **SC-003**: `go build ./...` 通过且无编译错误
- **SC-004**: `go vet ./...` 通过且无警告
- **SC-005**: 所有导出符号的 godoc 覆盖率达到 90%+
- **SC-006**: `go test ./...` 全部通过
- **SC-007**: `model/` 包测试覆盖率 ≥ 40%
- **SC-008**: 正常模型生成的代码与重构前完全一致
