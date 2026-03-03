# Feature Specification: 代码质量重构

**Feature Branch**: `001-code-quality-refactor`  
**Created**: 2026-03-03  
**Status**: Draft  
**Input**: 深入分析当前项目所有代码，找出需要优化改进的地方，让项目更优雅、可读性更强、更易于维护，并且适合大型项目架构。

## User Scenarios & Testing

### User Story 1 - 消除死代码和注释代码块 (Priority: P1)

作为开发者，我希望代码库中不存在被注释掉的代码块和死代码，以便在阅读代码时不会产生困惑，同时减少代码的认知负担。

**Why this priority**: 注释掉的代码块是代码库中最直接影响可读性的因素，分散注意力且增加维护成本，应最优先清理。

**Independent Test**: 通过全局搜索 `// if`、`// err =`、`// func` 等模式确认无残留注释代码块。

**Acceptance Scenarios**:

1. **Given** 代码库中存在多处 `// if err := ...`、`// func createDockerFile` 等注释代码块, **When** 执行重构, **Then** 所有注释代码块被移除，版本控制中保留历史记录
2. **Given** `cmd/gen.go` 中注释掉的 `model.BuildFederatedModel` 调用, **When** 执行重构, **Then** 被注释的代码块被彻底移除
3. **Given** `templates/helpers.go` 中注释掉的模板代码示例, **When** 执行重构, **Then** 示例代码被移除或转化为文档
4. **Given** `model/definition.inputs.go` 中大量注释代码块（第85-103行）, **When** 执行重构, **Then** 死代码被清除

---

### User Story 2 - 修复代码缺陷和逻辑错误 (Priority: P1)

作为开发者，我希望代码中不存在已知的 Bug 和逻辑错误，以保证代码生成器产出正确的代码。

**Why this priority**: 代码缺陷直接影响生成代码的正确性，是质量底线，需优先修复。

**Independent Test**: 为每个修复点编写单元测试，验证修复后行为正确。

**Acceptance Scenarios**:

1. **Given** `templates/helpers.go:176` 中 `isExist()` 函数使用 `os.IsExist(err)` 而非 `!os.IsNotExist(err)`, **When** 修复为正确实现, **Then** 函数在文件存在时返回 `true`，不存在时返回 `false`
2. **Given** `model/model.object-field-input.go:32-38` 中 `Required()` 方法使用 `bool` 作为变量名遮蔽内置类型且逻辑冗余, **When** 简化方法实现, **Then** 方法直接返回 `isNonNullType(o.Def.Type)` 的结果
3. **Given** `templates/utils-validator.go:160` 中 `Max()` 调用错误传入 `*minValue` 而非 `*maxValue`, **When** 修复为正确参数, **Then** 最大值校验使用正确的 `maxValue` 参数
4. **Given** `templates/utils-validator.go:124,137` 中 `Min()/Max()` 函数错误信息使用 `+` 拼接而非 `%s` 格式化（实际输出 "must be at least %s10"), **When** 修复格式化字符串, **Then** 错误信息正确显示为 "must be at least 10"
5. **Given** `model/printer.go:107` 中忽略 `printer.Print` 的类型断言错误, **When** 添加错误处理, **Then** 类型断言失败时返回有意义的错误信息

---

### User Story 3 - 消除重复代码和统一公共逻辑 (Priority: P2)

作为开发者，我希望代码库中不存在重复的功能实现，以降低维护成本并减少不一致变更的风险。

**Why this priority**: 重复代码是维护成本增长的主要来源，但不影响功能正确性。

**Independent Test**: 通过搜索确认所有公共函数只有唯一实现，调用方统一引用同一包。

**Acceptance Scenarios**:

1. **Given** `RunInteractiveInDir` 在 `templates/helpers.go:111` 和 `tools/run.go:12` 中存在两份重复实现（且行为不同：一个用 `sh`、一个用 `bash`）, **When** 统一为单一实现, **Then** 全项目只保留一份实现，调用方统一引用
2. **Given** `GetRandomString` 在 `model/utils.go:19` 和 `templates/resolver-utils.go:185` 中重复定义, **When** 合并为单一实现, **Then** 只保留一份且放在适当的包中
3. **Given** `IndexOf` 在 `model/utils.go:10` 和 `templates/resolver-utils.go:20` 中重复定义, **When** 合并重复代码, **Then** 统一为一处实现
4. **Given** `cmd/init.go` 中几乎每个 `create*File` 函数都重复调用 `model.LoadConfigFromPath(p)`, **When** 重构为一次加载多次使用, **Then** Config 只加载一次，通过参数传递

---

### User Story 4 - 使用废弃 API 升级 (Priority: P2)

作为开发者，我希望代码不使用已废弃的标准库 API，以保持与最新 Go 版本的兼容性。

**Why this priority**: 已废弃 API 可能在未来版本被移除，影响长期可维护性。

**Independent Test**: 运行 `go vet` 和 `staticcheck` 确认无废弃 API 警告。

**Acceptance Scenarios**:

1. **Given** `templates/helpers.go` 中使用已废弃的 `io/ioutil.WriteFile` 和 `ioutil.ReadFile`, **When** 替换为 `os.WriteFile` 和 `os.ReadFile`, **Then** 不再引入 `io/ioutil` 包
2. **Given** 文件写入使用 `0777` 权限模式, **When** 更改为更安全的 `0644`, **Then** 生成的文件权限更合理

---

### User Story 5 - 改进错误处理策略 (Priority: P2)

作为开发者，我希望代码使用一致的错误处理策略（返回 error 而非 panic），以便在运行时能优雅地处理异常场景。

**Why this priority**: panic 调用会导致程序崩溃，不利于生产环境稳定性。

**Independent Test**: 全局搜索确认 `panic(` 调用被替换为 `error` 返回，签名变更后调用方正确处理。

**Acceptance Scenarios**:

1. **Given** `model/model.go:96` 中 `Object()` 方法在对象不存在时 `panic`, **When** 修改为返回 `(Object, error)`, **Then** 调用方通过错误处理优雅应对
2. **Given** `model/model.go:105` 中 `ObjectExtension()` 方法类似地 `panic`, **When** 改为返回 error, **Then** 错误被传播而非崩溃
3. **Given** `model/model.object-relationship.go` 中多处 `panic`（第63、74、81行）, **When** 转换为 error 返回, **Then** 所有调用方正确处理错误
4. **Given** `model/utils.go:34` 中 `RegexpReplace` 忽略 `regexp.Compile` 的错误, **When** 添加错误处理, **Then** 正则表达式编译失败时返回有意义的错误
5. **Given** `templates/helpers.go:151` 中 `prompt` 函数在密码输入失败时 `panic`, **When** 改为返回 error, **Then** 密码读取失败时给出提示而非崩溃

---

### User Story 6 - 代码注释和文档规范化 (Priority: P3)

作为开发者，我希望代码中的注释遵循 Go 惯例（godoc 风格），且所有导出函数都有文档注释。

**Why this priority**: 规范化注释提升可读性和 IDE 体验，但不影响功能。

**Independent Test**: 运行 `golint` 或 `revive` 确认所有导出符号都有文档注释。

**Acceptance Scenarios**:

1. **Given** `cmd/root.go:9` 中 `// Execute ...` 这种非描述性注释, **When** 替换为规范的 godoc 注释, **Then** 注释清晰描述函数用途
2. **Given** `model/parser.go:7` 中 `// Parse` 这种空注释, **When** 补充完整注释, **Then** 注释描述了解析器的功能和输入输出
3. **Given** `model/enrichment.go:8-9` 中 `// EnrichModelObjects ...` 这种不完整注释, **When** 完善注释内容, **Then** 注释说明了函数如何丰富模型对象
4. **Given** 多个文件中混用中英文注释（如 `cmd/gen.go` 中 `// 接口` 和 `// 接口文档`）, **When** 统一注释语言风格, **Then** 中文注释保持一致格式

---

### User Story 7 - 模板文件结构优化 (Priority: P3)

作为开发者，我希望 Go 模板（字符串常量）被拆分为更小、更聚焦的单元，以降低单个文件的复杂度。

**Why this priority**: 改善代码组织结构，但需要谨慎操作以避免破坏现有功能。

**Independent Test**: 重构后运行 `make generate` 确认生成的代码与重构前完全一致（diff 比较）。

**Acceptance Scenarios**:

1. **Given** `templates/resolver-mutations.go` 是一个 784 行的单文件模板字符串, **When** 将辅助函数模板提取为独立常量, **Then** 每个模板专注于一个职责
2. **Given** `templates/resolver-queries.go` 是一个 301 行的模板字符串, **When** 拆分查询解析相关逻辑, **Then** 模板更易于理解和修改

---

### User Story 8 - 添加单元测试基础设施 (Priority: P3)

作为开发者，我希望项目有基本的单元测试覆盖关键逻辑，以便在重构时有安全网。

**Why this priority**: 测试是保证重构安全性的基础，但本身不改变产品功能。

**Independent Test**: 运行 `go test ./...` 确认所有测试通过。

**Acceptance Scenarios**:

1. **Given** 项目中不存在任何 `_test.go` 文件（已确认搜索结果为0）, **When** 为 `model/` 包的核心函数添加单元测试, **Then** 关键解析和类型转换逻辑有测试覆盖
2. **Given** `model/utils.go` 中的 `RegexpReplace` 和 `GetRandomString` 无测试, **When** 添加单元测试, **Then** 边界条件和常规用例都有测试覆盖
3. **Given** `utils/file_utils.go` 中的 `EnsureDir` 和 `FileExists` 无测试, **When** 添加测试, **Then** 文件存在/不存在/权限异常等场景都被覆盖

---

### Edge Cases

- 重构后的 `generate` 命令生成的代码是否与重构前完全一致？
- 修改 `panic` 为 `error` 返回后，所有调用方是否正确传播错误？
- 统一 `RunInteractiveInDir` 实现时，选择 `sh` 还是 `bash` 是否会导致跨平台兼容性问题？
- `isExist()` 函数修复后，是否有调用方依赖错误的旧行为？

## Requirements

### Functional Requirements

- **FR-001**: 代码库不得包含被注释掉的代码块（版本控制已保存历史）
- **FR-002**: 所有导出函数必须使用 Go error 返回模式代替 panic
- **FR-003**: 公共工具函数（`RunInteractiveInDir`、`GetRandomString`、`IndexOf`）在全项目中必须只有唯一实现
- **FR-004**: 代码不得使用 Go 标准库已废弃的 API（如 `io/ioutil`）
- **FR-005**: `templates/helpers.go` 中的 `isExist()` 函数必须返回正确的文件存在性判断结果
- **FR-006**: `model/model.object-field-input.go` 中的 `Required()` 方法必须简化为直接返回类型判断结果
- **FR-007**: `cmd/init.go` 中的配置加载必须统一为单次加载、参数传递模式
- **FR-008**: `templates/utils-validator.go` 中的 `validateNumberRange` 必须使用正确的 `maxValue` 参数
- **FR-009**: 所有导出符号必须有符合 godoc 规范的文档注释
- **FR-010**: `model/` 包的核心函数必须有单元测试覆盖
- **FR-011**: 生成文件的权限模式必须从 `0777` 改为 `0644`
- **FR-012**: `templates/utils-validator.go` 中的 `Min()/Max()` 函数错误信息格式必须正确

### Assumptions

- 项目使用 Git 进行版本控制，注释代码块可以安全删除（历史记录可追溯）
- 重构不改变代码生成器的输出结果（行为等价性）
- Go 1.24+ 环境，支持泛型和最新标准库
- `gqlgen` 依赖（v0.17.85）为锁定版本，不在本次重构范围内
- 模板字符串的结构优化可以分阶段进行，不必一次完成

## Success Criteria

### Measurable Outcomes

- **SC-001**: 代码库中被注释掉的代码块数量从当前 15+ 处降至 0 处
- **SC-002**: `panic` 调用数量从当前 8+ 处降至 0 处（测试中的 panic 除外）
- **SC-003**: 重复函数实现数量从当前 3 组降至 0 组
- **SC-004**: `go vet ./...` 执行通过且无警告
- **SC-005**: `model/` 包单元测试覆盖率达到 60% 以上
- **SC-006**: 重构后 `make generate` 生成的代码与重构前完全一致（通过 diff 验证）
- **SC-007**: 所有导出符号的 godoc 覆盖率达到 90% 以上
- **SC-008**: 不使用任何已废弃的标准库 API
