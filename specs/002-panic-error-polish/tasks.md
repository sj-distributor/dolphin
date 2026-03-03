# Tasks: Panic→Error 重构与代码完善

**Input**: Design documents from `/specs/002-panic-error-polish/`
**Prerequisites**: plan.md, spec.md, research.md

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)

---

## Phase 1: Setup

**Purpose**: 确认基线编译状态

- [x] T001 确认项目可成功编译：`go build ./...`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: 无（所有变更直接在 User Story 中）

**Checkpoint**: 直接进入 User Story 实施

---

## Phase 3: User Story 1 - panic → error 返回 (Priority: P1) 🎯 MVP

**Goal**: 将 6 处 panic 替换为 error 返回，确保错误沿调用链传播至 CLI 层

**Independent Test**: 构造引用不存在类型的 `.graphql` 文件，运行 `dolphin generate`，确认输出错误信息（非崩溃）

### Phase 3a: 底层方法签名变更

- [x] T002 [US1] 修改 `model/model.object-relationship.go` 中 `StringForRelationshipDirectiveAttribute()` 签名为 `(string, error)`，将第63行 panic 替换为 `return "", fmt.Errorf(...)`
- [x] T003 [US1] 修改 `model/model.object-relationship.go` 中 `BoolForRelationshipDirectiveAttribute()` 签名为 `(bool, error)`，将第74行 panic 替换为 `return false, fmt.Errorf(...)`
- [x] T004 [US1] 修改 `model/model.object-relationship.go` 中 `InverseRelationshipName()` 签名为 `(string, error)`，将第81行 panic 替换为 `return "", fmt.Errorf(...)`
- [x] T005 [US1] 更新 `model/model.object-relationship.go` 内部调用方 `Target()`, `InverseRelationship()`, `JoinTable()`, `IsManyToMany()`, `IsOneToOne()`, `IsManyToOne()` 等方法以处理新的 error 返回
- [x] T006 [US1] 运行 `go build ./model/...` 确认 model.object-relationship.go 编译通过

### Phase 3b: 核心模型方法签名变更

- [x] T007 [P] [US1] 修改 `model/model.go` 中 `Object(name)` 签名为 `(Object, error)`，将第96行 panic 替换为 `return Object{}, fmt.Errorf(...)`
- [x] T008 [P] [US1] 修改 `model/model.go` 中 `ObjectExtension(name)` 签名为 `(ObjectExtension, error)`，将第105行 panic 替换为 `return ObjectExtension{}, fmt.Errorf(...)`
- [x] T009 [P] [US1] 修改 `model/model.object.go` 中 `Relationship(name)` 签名为 `(ObjectRelationship, error)`，将第157行 panic 替换为 `return ObjectRelationship{}, fmt.Errorf(...)`
- [x] T010 [P] [US1] 修改 `model/config.go` 中 `ConnMaxLifetime()` 签名为 `(time.Duration, error)`，将第57行 panic 替换为 `return 0, fmt.Errorf(...)`

### Phase 3c: 调用方更新（错误传播）

- [x] T011 [US1] 更新 `model/model.object-field.go` 中所有调用 `Object()`, `ObjectExtension()`, `Relationship()` 的方法以处理 error 返回：`TargetObject()`, `TargetObjectExtension()`, `HasTargetObject()`, `HasTargetObjectExtension()`, `HasTargetTypeWithIDField()`
- [x] T012 [US1] 更新 `model/enrichment.go` 中 `EnrichModelObjects()` 调用 `Object()` 和 relationship 方法的地方以处理 error 传播
- [x] T013 [P] [US1] 修改 `model/utils.go` 中 `RegexpReplace()` 添加 `regexp.Compile` 错误处理（编译失败返回原字符串）
- [x] T014 [P] [US1] 修改 `templates/helpers.go` 中 `prompt()` 函数的密码读取 panic 替换为 error 返回，更新 `Prompt()` 调用方
- [x] T015 [US1] 确认 `cmd/gen.go` 中 `EnrichModel()` 的 error 已正确传播到 `cli.NewExitError()`
- [x] T016 [US1] 运行 `go build ./...` 确认全项目编译通过
- [x] T017 [US1] 运行 `go vet ./...` 确认零警告

**Checkpoint**: 所有 panic 已替换，错误正确传播，编译通过

---

## Phase 4: User Story 2 - 代码注释规范化 (Priority: P2)

**Goal**: 所有导出符号具有 godoc 规范注释

**Independent Test**: 无 "exported function XXX should have comment" 类警告

### Implementation

- [x] T018 [P] [US2] 规范化 `cmd/root.go` 中 `Execute` 函数的 godoc 注释
- [ ] T019 [P] [US2] 规范化 `model/parser.go` 中 `Parse` 函数的 godoc 注释
- [ ] T020 [P] [US2] 规范化 `model/enrichment.go` 中 `EnrichModelObjects` 和 `EnrichModel` 的 godoc 注释
- [ ] T021 [P] [US2] 规范化 `model/model.go` 中所有导出方法的 godoc 注释（`SecretKey`, `Objects`, `HasObject`, `Object`, `ObjectExtension`, `ObjectEntities`）
- [ ] T022 [P] [US2] 规范化 `model/model.object.go` 中所有导出方法的 godoc 注释
- [ ] T023 [P] [US2] 规范化 `model/graphql-helpers.go` 中所有函数的 godoc 注释
- [ ] T024 [P] [US2] 规范化 `templates/helpers.go` 中 `WriteTemplate`, `WriteTemplateRaw`, `WriterOriginalFile`, `WriteInterfaceTemplate` 的 godoc 注释
- [ ] T025 [P] [US2] 规范化 `model/definition.*.go` 系列文件中的注释风格

**Checkpoint**: 所有导出符号有规范注释

---

## Phase 5: User Story 3 - 单元测试 (Priority: P3)

**Goal**: 核心函数有单元测试覆盖，`model/` 包覆盖率 ≥ 40%

**Independent Test**: `go test ./model/ ./utils/` 全部通过

### Implementation

- [x] T026 [P] [US3] 创建 `model/utils_test.go`：测试 `GetRandomString`（长度、字符范围）、`RegexpReplace`（正常替换、无效正则返回原字符串）、`IndexOf`（找到/未找到）
- [x] T027 [P] [US3] 创建 `utils/file_utils_test.go`：测试 `EnsureDir`（新建/已存在）、`FileExists`（存在/不存在/目录）
- [ ] T028 [P] [US3] 创建 `model/model_test.go`：测试 `Parse`（基本 GraphQL 解析）、`Object`（找到/未找到返回 error）、`ObjectEntities`（过滤非实体）
- [x] T029 [US3] 运行 `go test ./model/ ./utils/` 确认所有测试通过
- [ ] T030 [US3] 运行 `go test -cover ./model/` 确认覆盖率 ≥ 40%

**Checkpoint**: 核心函数有测试覆盖

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: 最终验证

- [x] T031 运行 `go build ./...` 最终编译验证
- [x] T032 运行 `go vet ./...` 确认零警告
- [x] T033 运行 `go test ./...` 全部通过
- [ ] T034 在 `example/` 运行 `go run . generate` 确认正常模型生成代码无变化

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: 无依赖
- **US1 panic→error (Phase 3)**: 依赖 Setup — **必须按 3a → 3b → 3c 顺序**（从底向上修改）
- **US2 注释规范化 (Phase 4)**: 依赖 US1（函数签名稳定后再写注释）
- **US3 单元测试 (Phase 5)**: 依赖 US1（API 稳定后再写测试）
- **Polish (Phase 6)**: 依赖所有 US 完成

### Within US1 (Critical Path)

```
T002-T004 (底层 panic 替换) → T005 (内部调用方) → T006 (编译检查)
  → T007-T010 (核心方法签名变更，可并行) → T011-T012 (调用方更新)
  → T013-T014 (独立修复，可并行) → T015-T017 (最终验证)
```

### Parallel Opportunities

- **T007-T010**: 4 个核心方法签名变更可并行（不同文件）
- **T013-T014**: utils.go 和 helpers.go 修改可并行
- **T018-T025**: 所有注释规范化任务可完全并行
- **T026-T028**: 所有测试文件创建可完全并行

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 3: US1 panic→error（按 3a → 3b → 3c 严格顺序）
3. **STOP and VALIDATE**: 编译通过 + 行为等价性验证
4. 这是最核心改进，其他 US 可增量交付

### Incremental Delivery

1. Setup → US1 panic→error → **验证** (MVP)
2. + US2 注释规范化 → **验证**
3. + US3 单元测试 → **验证**
4. Polish → **最终验证**

---

## Notes

- US1 中签名变更必须**严格从底向上**执行，否则中间状态无法编译
- `getNamedType()` 的 panic 保留（编程错误），不在任务范围内
- 模板文件（`templates/*.go`）中的字符串模板不受方法签名变更影响
- 错误消息使用 `fmt.Errorf` 格式化，包含足够上下文（对象名、关系名等）
