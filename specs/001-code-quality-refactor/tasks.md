# Tasks: 代码质量重构

**Input**: Design documents from `/specs/001-code-quality-refactor/`
**Prerequisites**: plan.md, spec.md, research.md

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)

---

## Phase 1: Setup

**Purpose**: 创建验证基准（重构前生成代码的快照）

- [x] T001 在 `master` 分支上对 `example/` 运行 `go run . generate`，将 `example/gen/` 和 `example/utils/` 备份到 `/tmp/dolphin-before/`
- [x] T002 切回 `001-code-quality-refactor` 分支，确认项目可成功编译：`go build ./...`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: 无（本特性无共享基础设施依赖）

**Checkpoint**: 直接进入 User Story 实施

---

## Phase 3: User Story 1 - 消除死代码和注释代码块 (Priority: P1) 🎯 MVP

**Goal**: 移除所有被注释掉的代码块，提升代码可读性

**Independent Test**: 全局搜索 `// if`、`// err =`、`// func` 等模式确认零残留

### Implementation

- [x] T003 [P] [US1] 移除 `cmd/gen.go` 中注释掉的 `model.BuildFederatedModel` 调用（第123-126行）和 `ResolverSrcContext`（第233-235行）
- [x] T004 [P] [US1] 移除 `cmd/init.go` 中注释掉的 `createDockerFile` 调用（第72-74行）、函数体（第162-169行）、`UploadModel`（第144-146行）和 `UpLoad`（第233-235行）
- [x] T005 [P] [US1] 移除 `model/definition.inputs.go` 中注释掉的旧条件逻辑代码块（第85-103行）
- [x] T006 [P] [US1] 移除 `model/definition.filter.go` 中注释的 id filter（第18行）和 Description（第52-55行）
- [x] T007 [P] [US1] 移除 `model/definition.mutation.go` 中注释掉的 Description 字段（第38、62、95、100、112行）
- [x] T008 [P] [US1] 移除 `model/definition.query.go` 中注释掉的代码（第12、46、53、60、100行）
- [x] T009 [P] [US1] 移除 `model/enrichment.go` 中注释掉的 `columnDefinition`（第26行）和 `createFederationServiceObject`（第65行）
- [x] T010 [P] [US1] 移除 `templates/helpers.go` 中注释掉的示例代码（第42-52行）和 `set -o pipefail`（第116行）
- [ ] T011 [P] [US1] 移除 `templates/resolver-utils.go` 模板中注释掉的 EventController 代码块（第74-77行）
- [ ] T012 [P] [US1] 移除 `templates/resolver-queries.go` 模板中注释掉的 select/where 查询代码块（第213-217行）
- [x] T013 [US1] 运行 `go build ./...` 确认编译通过

**Checkpoint**: 所有注释代码块已清除，项目编译正常

---

## Phase 4: User Story 2 - 修复代码缺陷和逻辑错误 (Priority: P1)

**Goal**: 修复所有已知 Bug，保证代码生成器产出正确代码

**Independent Test**: 对每个修复点手动验证修复后行为

### Implementation

- [x] T014 [P] [US2] 修复 `templates/helpers.go` 中 `isExist()` 函数：`os.IsExist(err)` → `err == nil`（第174-177行）
- [x] T015 [P] [US2] 简化 `model/model.object-field-input.go` 中 `Required()` 方法：移除冗余的 `isEmpty`/`bool` 变量，直接返回 `isNonNullType(o.Def.Type)`（第32-38行）
- [x] T016 [P] [US2] 修复 `templates/utils-validator.go` 中 `validateNumberRange` 的 `Max()` 调用参数错误：`Max(intValue, *minValue)` → `Max(intValue, *maxValue)`（第160行）
- [x] T017 [P] [US2] 修复 `templates/utils-validator.go` 中 `Min()` 格式化字符串错误：`"must be at least %s" + convertor.ToString(min)` → `fmt.Sprintf("must be at least %s", convertor.ToString(min))`（第124行）
- [x] T018 [P] [US2] 修复 `templates/utils-validator.go` 中 `Max()` 格式化字符串错误：同 T017 修复模式（第137行）
- [x] T019 [P] [US2] 修复 `model/printer.go` 中 `PrintSchema` 函数忽略 `printer.Print` 类型断言错误（第107行）
- [x] T020 [US2] 运行 `go build ./...` 确认编译通过

**Checkpoint**: 所有已知 Bug 已修复

---

## Phase 5: User Story 3 - 消除重复代码和统一公共逻辑 (Priority: P2)

**Goal**: 消除全项目中重复的函数实现，统一为单一入口

**Independent Test**: 全局搜索确认每个公共函数只有唯一实现

### Implementation

- [x] T021 [US3] 重构 `cmd/init.go`：在 `initCmd.Action` 顶层一次性加载 Config，修改所有 `create*File` 函数签名接受 `*model.Config` 参数，移除每个函数内重复的 `model.LoadConfigFromPath` 调用
- [x] T022 [US3] 统一 `RunInteractiveInDir` 实现：保留 `templates/helpers.go` 中的版本，修改 `cmd/gen.go` 第159行从 `tools.RunInteractiveInDir` 改为 `templates.RunInteractiveInDir`，移除 `tools` import
- [x] T023 [US3] 删除 `tools/run.go` 文件（确认无其他引用后）
- [x] T024 [US3] 运行 `go build ./...` 确认编译通过

**Checkpoint**: 全项目无重复函数实现

---

## Phase 6: User Story 4 - 废弃 API 升级 (Priority: P2)

**Goal**: 移除所有废弃标准库 API 的使用

**Independent Test**: 运行 `go vet ./...` 确认无废弃 API 警告

### Implementation

- [x] T025 [US4] 修改 `templates/helpers.go`：替换 `ioutil.WriteFile` → `os.WriteFile`，`ioutil.ReadFile` → `os.ReadFile`，移除 `io/ioutil` import，文件权限 `0777` → `0644`
- [x] T026 [US4] 运行 `go vet ./...` 确认无警告

**Checkpoint**: 无废弃 API 使用

---

## Phase 7: User Story 5 - 改进错误处理策略 (Priority: P2)

**Goal**: 将所有 panic 调用替换为 error 返回模式

**Independent Test**: 全局搜索确认非测试代码中 `panic(` 调用为 0

### Implementation

- [ ] T027 [US5] 修改 `model/model.go`：将 `Object(name)` 方法签名从 `Object` 改为 `(Object, error)`，将 `ObjectExtension(name)` 从 `ObjectExtension` 改为 `(ObjectExtension, error)`，替换 panic 为 error 返回
- [ ] T028 [US5] 更新 `model/model.go` 中 `Object()` 和 `ObjectExtension()` 的所有调用方（`model/enrichment.go`、`model/model.object-relationship.go`、`model/model.object-field.go` 等）以处理新的 error 返回
- [ ] T029 [US5] 修改 `model/model.object-relationship.go`：将 `StringForRelationshipDirectiveAttribute`（第63行）、`BoolForRelationshipDirectiveAttribute`（第74行）和 `InverseRelationshipName`（第81行）中的 panic 替换为 error 返回，更新调用方
- [ ] T030 [US5] 修改 `model/config.go`：将 `ConnMaxLifetime()` 中的 panic（第57行）替换为 error 返回，更新调用方
- [ ] T031 [US5] 修改 `model/utils.go`：为 `RegexpReplace` 中的 `regexp.Compile`（第34行）添加错误处理
- [ ] T032 [US5] 修改 `templates/helpers.go`：将 `prompt` 函数中密码读取的 panic（第151行）替换为 error 返回，更新调用方 `Prompt()`
- [ ] T033 [US5] 运行 `go build ./...` 确认编译通过，全局搜索确认 `panic(` 在非测试代码中为 0

**Checkpoint**: 所有 panic 已替换为优雅的错误处理

---

## Phase 8: User Story 6 - 代码注释和文档规范化 (Priority: P3)

**Goal**: 所有导出符号具有 godoc 规范注释

**Independent Test**: 运行 linter 确认所有导出符号有文档注释

### Implementation

- [ ] T034 [P] [US6] 规范化 `cmd/root.go` 中 `Execute` 函数的注释
- [ ] T035 [P] [US6] 规范化 `model/parser.go` 中 `Parse` 函数的注释
- [ ] T036 [P] [US6] 规范化 `model/enrichment.go` 中 `EnrichModelObjects` 和 `EnrichModel` 函数的注释
- [ ] T037 [P] [US6] 规范化 `model/model.go` 中所有导出方法的注释（`SecretKey`、`Objects`、`HasObject` 等）
- [ ] T038 [P] [US6] 规范化 `model/model.object.go` 中所有导出方法的注释
- [ ] T039 [P] [US6] 规范化 `model/graphql-helpers.go` 中所有函数的注释
- [ ] T040 [P] [US6] 规范化 `templates/helpers.go` 中 `WriteTemplate`、`WriteTemplateRaw`、`WriterOriginalFile` 等导出函数的注释
- [ ] T041 [P] [US6] 统一其余 `model/definition.*.go` 文件中的注释风格

**Checkpoint**: 所有导出符号有规范注释

---

## Phase 9: User Story 8 - 添加单元测试基础设施 (Priority: P3)

**Goal**: 为核心逻辑添加单元测试

**Independent Test**: 运行 `go test ./...` 全部通过

### Implementation

- [ ] T042 [P] [US8] 创建 `model/model_test.go`：测试 `Parse()`、`Objects()`、`HasObject()`、`ObjectEntities()` 等核心方法
- [ ] T043 [P] [US8] 创建 `model/utils_test.go`：测试 `GetRandomString()`、`RegexpReplace()`、`IndexOf()`
- [ ] T044 [P] [US8] 创建 `model/object_field_test.go`：测试 `GoTypeWithPointer()`、`IsCreatable()`、`IsUpdatable()`、`FilterMapping()`
- [ ] T045 [P] [US8] 创建 `utils/file_utils_test.go`：测试 `EnsureDir()`、`FileExists()`
- [ ] T046 [US8] 运行 `go test ./...` 确认所有测试通过

**Checkpoint**: 核心逻辑有基础测试覆盖

---

## Phase 10: Polish & Cross-Cutting Concerns

**Purpose**: 最终验证和清理

- [x] T047 对 `example/` 运行 `go run . generate`，将输出与 `/tmp/dolphin-before/` 进行 diff 比对，确认行为等价性（Bug 修复导致的差异需标注为预期变更）
- [x] T048 运行 `go vet ./...` 确认零警告
- [ ] T049 运行 `go test ./...` 确认所有测试通过
- [x] T050 运行 `go build ./...` 最终编译验证

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: 无依赖 — 立即开始
- **US1 死代码清除 (Phase 3)**: 依赖 Setup — 其他 US 可独立于 US1
- **US2 Bug 修复 (Phase 4)**: 依赖 Setup — 可与 US1 并行
- **US3 重复代码 (Phase 5)**: 依赖 US1（清除死代码后再重构更清晰）
- **US4 废弃 API (Phase 6)**: 依赖 US3（helpers.go 重构完成后）
- **US5 错误处理 (Phase 7)**: 依赖 US1 + US3（代码清理后再改签名更安全）
- **US6 注释规范化 (Phase 8)**: 依赖 US5（函数签名稳定后再写注释）
- **US8 单元测试 (Phase 9)**: 依赖 US5（API 稳定后再写测试）
- **Polish (Phase 10)**: 依赖所有 US 完成

### Parallel Opportunities

- **T003-T012**: 所有 US1 死代码清除任务可完全并行
- **T014-T019**: 所有 US2 Bug 修复任务可完全并行
- **T034-T041**: 所有 US6 注释规范化任务可完全并行
- **T042-T045**: 所有 US8 单元测试创建任务可完全并行

---

## Parallel Example: User Story 1

```bash
# 所有死代码清除任务可同时执行（不同文件，无依赖）：
Task T003: "移除 cmd/gen.go 中注释代码"
Task T004: "移除 cmd/init.go 中注释代码"
Task T005: "移除 model/definition.inputs.go 中注释代码"
Task T006: "移除 model/definition.filter.go 中注释代码"
Task T007: "移除 model/definition.mutation.go 中注释代码"
Task T008: "移除 model/definition.query.go 中注释代码"
Task T009: "移除 model/enrichment.go 中注释代码"
Task T010: "移除 templates/helpers.go 中注释代码"
Task T011: "移除 templates/resolver-utils.go 中注释代码"
Task T012: "移除 templates/resolver-queries.go 中注释代码"
```

---

## Implementation Strategy

### MVP First (User Story 1 + 2 Only)

1. Complete Phase 1: Setup（创建基准快照）
2. Complete Phase 3: US1 死代码清除 + Phase 4: US2 Bug 修复
3. **STOP and VALIDATE**: 编译通过 + 行为等价性比对
4. 这两个 P1 故事构成最小可行重构

### Incremental Delivery

1. Setup → US1 + US2 → **验证** (MVP)
2. + US3 重复代码 + US4 废弃 API → **验证**
3. + US5 错误处理 → **验证**
4. + US6 注释 + US8 测试 → **验证**
5. Polish → **最终验证**

---

## Notes

- US7（模板文件结构优化）已从任务列表中排除 — 风险较高且收益有限，建议作为独立后续特性
- 每个 Phase Checkpoint 后都应提交代码（atomic commits）
- 修改函数签名（US5）可能影响范围较广，需特别注意调用链
- 模板代码（`templates/*.go`）中的"重复"是设计意图（生成到目标项目），不应消除
