# Tasks: Enum API Metadata

**Input**: Design documents from `/specs/006-enum-api-metadata/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, quickstart.md

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 [P] Ensure `example/model/test.graphql` has enums with `@entity` directives

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

- [x] T002 Verify `GetDefinition(name string)` in `model/model.go` handles all enum definitions
- [x] T003 [P] Add `getEnumTitle(v *ast.EnumValueDefinition)` helper in `model/model.object-field.go`

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - See Enum Value Titles in Documentation (Priority: P1) 🎯 MVP

**Goal**: Extract and display enum value titles in `api.json`

**Independent Test**: Run `make generate` in `example/` and check `docs/api.json` for enum titles

### Implementation for User Story 1

- [x] T004 [US1] Implement `Case *ast.EnumDefinition` in `GetTypeData` within `model/model.object-field.go`
- [x] T005 [US1] Verify enum metadata in `example/docs/api.json` after generation

**Checkpoint**: At this point, User Story 1 should be fully functional

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T006 [P] Run regression tests with `go test ./model/...`
- [x] T007 [P] Documentation updates in `walkthrough.md`
- [x] T008 Run `quickstart.md` validation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies
- **Foundational (Phase 2)**: Depends on Setup completion
- **User Story 1 (Phase 3)**: Depends on Foundational completion
- **Polish (Final Phase)**: Depends on story completion

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Verify enum titles in `api.json`
