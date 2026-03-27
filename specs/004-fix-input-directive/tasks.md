# Tasks: Fix Custom Input Directive Processing

**Input**: Design documents from `/specs/004-fix-input-directive/`
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

- [x] T002 Analyze `model/printer.go` for all directive-bearing AST node types
- [x] T003 [P] Implement `filterDirective(ds []*ast.Directive, name string)` helper in `model/printer.go` (existing, but check for adequacy)

## Phase 3: User Story 1 - Support Custom Query with Input Directives (Priority: P1) 🎯 MVP

**Goal**: Enable `make generate` to succeed when custom queries use `input` types with directives.

**Independent Test**: Define a custom `Query` in `example/model/test.graphql` using an `input` type with `@entity` directives, and run `make generate` successfully.

### Implementation for User Story 1

- [x] T004 [US1] Implement `cleanNode(node interface{})` or similar recursion in `model/printer.go` to handle `ObjectDefinition` and `InputObjectDefinition`
- [x] T005 [US1] Ensure `cleanNode` handles `FieldDefinition` and its arguments (`InputValueDefinition`) recursively
- [x] T006 [US1] Update `PrintSchema` in `model/printer.go` to call `cleanNode` for all definitions in `model.Doc`
- [x] T007 [US1] Verify fix by running `make generate` in `example/`
- [x] T008 [US1] Check `example/gen/schema.graphqls` to ensure all dolphin directives are removed

## Phase N: Polish & Cross-Cutting Concerns

- [x] T009 [P] Run overall tests with `go test ./model/...`
- [x] T010 [P] Code cleanup and redundant code removal in `model/printer.go` (remove old loops)
- [x] T011 Run `quickstart.md` validation steps

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: Completed
- **Foundational (Phase 2)**: BLOCKS user story implementation
- **User Story 1 (P1)**: Depends on Foundational completion
- **Polish (Final Phase)**: Depends on US1 completion

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1 & 2
2. Complete Phase 3: User Story 1
3. **STOP and VALIDATE**: Test `make generate` in `example/`
