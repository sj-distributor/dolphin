# Tasks: Extension API Metadata

**Input**: Design documents from `/specs/005-extension-api-metadata/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize feature branch and environment

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

- [x] T002 Implement `GetDefinition(name string) ast.Node` in `model/model.go`
- [x] T003 [P] Define `TypeData` and `FieldMetadata` structs in `model/model.go`

## Phase 3: User Story 1 - See Custom API Metadata in Documentation (Priority: P1) 🎯 MVP

**Goal**: Populate `typeData` in `api.json` for custom extensions.

**Independent Test**: Run `make generate` in `example/` and verify `api.json` has populated `typeData` for `login`.

### Implementation for User Story 1

- [x] T004 [US1] Implement `GetTypeData() []TypeData` in `model/model.object-field.go`
- [x] T005 [P] [US1] Update `templates/graphql.go` to iterate over `GetTypeData()` and populate the `typeData` JSON field
- [x] T006 [US1] Verify fix by running `make generate` in `example/` and checking `example/docs/api.json`

## Phase N: Polish & Cross-Cutting Concerns

- [x] T007 [P] Run overall tests with `go test ./model/...`
- [x] T008 [P] Documentation updates and code cleanup
- [x] T009 Run `quickstart.md` validation steps

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Completed
- **Foundational (Phase 2)**: BLOCKS user story implementation
- **User Story 1 (P1)**: Depends on Foundational completion
- **Polish (Final Phase)**: Depends on US1 completion

## Parallel Execution Examples

```bash
# Foundational
Task: T002 Implement GetDefinition
Task: T003 Define structs

# Implementation
Task: T004 Implement GetTypeData
Task: T005 Update template
```

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1 & 2
2. Complete Phase 3: User Story 1
3. **STOP and VALIDATE**: Test `make generate` in `example/`
