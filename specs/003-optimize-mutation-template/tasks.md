# Tasks: Mutation Template Optimization

**Input**: Design documents from `/specs/003-optimize-mutation-template/`
**Prerequisites**: plan.md (7 changes), spec.md (7 user stories), research.md (9 decisions)

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)

---

## Phase 1: Setup

**Purpose**: Confirm baseline compilation state before making changes

- [x] T001 Confirm project compiles successfully: `go build ./...` in `/Users/marlon.m/wwwroot/triton/dolphin/`
- [x] T002 Confirm existing tests pass: `go test ./model/ ./utils/` in `/Users/marlon.m/wwwroot/triton/dolphin/`
- [x] T003 Verify `$rel.InverseRelationship` template variable exists in `model/` package — needed for US7 OneToOne unbinding implementation

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: None — all changes are in a single template file with no shared foundational dependencies.

---

## Phase 3: US1 — Error Message Formatting (Priority: P1) 🎯 MVP

**Goal**: Replace all string concatenation error patterns with `fmt.Errorf` + `%w` wrapping  
**Independent Test**: `grep 'errors.New.*err.Error' templates/resolver-mutations.go` returns 0 matches

### Implementation

- [x] T004 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Create handler ToOne section (L118-139) in `templates/resolver-mutations.go`
- [x] T005 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Create handler ToMany IDs section (L215-230) in `templates/resolver-mutations.go`
- [x] T006 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Create handler ToMany nested section (L264-288) in `templates/resolver-mutations.go`
- [x] T007 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Update handler ToOne section (L420-445) in `templates/resolver-mutations.go`
- [x] T008 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Update handler ToMany IDs section (L529-541) in `templates/resolver-mutations.go`
- [x] T009 [US1] Replace `errors.New(...)` with `fmt.Errorf("...: %w", err)` in Update handler ToMany nested section (L588-610) in `templates/resolver-mutations.go`
- [x] T010 [US1] Replace `fmt.Errorf("..." + err.Error())` with `fmt.Errorf("...: %w", err)` in Create/Update handler column sections (L179, L484) in `templates/resolver-mutations.go`
- [x] T011 [US1] Run `go build ./...` to confirm compilation passes

**Checkpoint**: Zero instances of `errors.New("..." + err.Error())` or string concatenation in error messages

---

## Phase 4: US2 — UpdatedAt/UpdatedBy Fix (Priority: P1) 🎯 MVP

**Goal**: Ensure `updated_at` and `updated_by` are always set on every Update mutation  
**Independent Test**: Generated Update handler contains `updated_at` in `changedFields`

### Implementation

- [x] T012 [US2] Replace conditional `UpdatedBy` logic (L399-402) with unconditional `newItem.UpdatedAt = &timestampMillis` and `newItem.UpdatedBy = principalID` in `templates/resolver-mutations.go`
- [x] T013 [US2] Update `changedFields` append (L503-506): replace conditional `updated_by` with `changedFields = append(changedFields, "updated_at", "updated_by")` in `templates/resolver-mutations.go`
- [x] T014 [US2] Run `go build ./...` to confirm compilation passes

**Checkpoint**: Update handler always includes `updated_at` and `updated_by` in SQL UPDATE

---

## Phase 5: US6 — OneToMany Nested Update Old Association Cleanup (Priority: P1) 🆕

**Goal**: Clear orphaned OneToMany associations before processing nested objects in Update handler  
**Independent Test**: Generated Update handler's ToMany nested section contains `UPDATE SET fk = NULL` before the `for` loop (for non-ManyToMany relationships)

### Implementation

- [x] T015 [US6] Add OneToMany old association clearing block before the nested object `for` loop (~L575) in Update handler ToMany section: `tx.Model(&{{TargetType}}{}).Where("fk = ?", item.ID).Update("fk", nil)` — only for `{{if not $rel.IsManyToMany}}` in `templates/resolver-mutations.go`
- [x] T016 [US6] Run `go build ./...` to confirm compilation passes

**Checkpoint**: OneToMany nested update clears old associations not in submitted list

---

## Phase 6: US7 — OneToOne Update Unbind Old Association (Priority: P2) 🆕

**Goal**: Unbind old OneToOne associated object before binding new one in Update handler  
**Independent Test**: Generated Update handler's ToOne section contains reverse FK clearing logic before FK update

### Implementation

- [x] T017 [US7] Add old OneToOne association unbinding block before the existing update/create branches (~L412) in Update handler ToOne section — query current FK, set old object's reverse FK to nil. Guard with `{{if $rel.InverseRelationship}}` in `templates/resolver-mutations.go`
- [x] T018 [US7] Run `go build ./...` to confirm compilation passes

**Checkpoint**: OneToOne update properly unbinds old associated object

---

## Phase 7: US3 — Permission Check Optimization (Priority: P2)

**Goal**: Execute permission checks once per relationship type, not once per nested object  
**Independent Test**: Generated code shows `CheckAuthorization` call before/outside `for` loop with flag guard

### Implementation

- [x] T019 [US3] Add `hasCreate`/`hasUpdate` flag guards in Create handler ToMany nested loop (~L255-310) in `templates/resolver-mutations.go`
- [x] T020 [US3] Add `hasCreate`/`hasUpdate` flag guards in Update handler ToMany nested loop (~L579-631) in `templates/resolver-mutations.go`
- [x] T021 [US3] Run `go build ./...` to confirm compilation passes

**Checkpoint**: Permission checks execute O(1) per relationship type

---

## Phase 8: US4 — Create ToOne Conflict Detection (Priority: P2)

**Goal**: Add ToOne relationship conflict validation to Create handler  
**Independent Test**: Generated Create handler's validation block contains `cannot coexist` for both ToMany and ToOne relationships

### Implementation

- [x] T022 [US4] Add `{{else}}` branch with ToOne conflict check `fmt.Errorf("{{rel.Name}}Id and {{rel.Name}} cannot coexist")` in Create handler validation block (~L96-103) in `templates/resolver-mutations.go`
- [x] T023 [US4] Run `go build ./...` to confirm compilation passes

**Checkpoint**: Create and Update have symmetric conflict validation

---

## Phase 9: US5 — Create Change Detection Simplification (Priority: P3)

**Goal**: Remove redundant zero-value comparison in Create handler field assignment  
**Independent Test**: Generated Create handler's column section uses direct assignment without `item.X != changes.X` guard

### Implementation

- [x] T024 [US5] Simplify Create handler column assignment block (~L168-191): remove outer `if (item.{{$col.MethodName}} != changes.{{$col.MethodName}})` comparison and the `!utils.IsEmpty()` guard for non-optional fields in `templates/resolver-mutations.go`
- [x] T025 [US5] Run `go build ./...` to confirm compilation passes

**Checkpoint**: Create handler uses cleaner direct assignment

---

## Phase 10: Polish & Cross-Cutting Concerns

**Purpose**: Final verification of all changes

- [x] T026 Run `go build ./...` final compilation check
- [x] T027 Run `go vet ./...` confirm zero warnings
- [x] T028 Run `go test ./model/ ./utils/` confirm all tests pass
- [x] T029 Grep template for remaining `+ err.Error()` patterns — confirm 0 matches in `templates/resolver-mutations.go`
- [x] T030 Grep template for remaining `errors.New("` patterns — confirm 0 matches (except legitimate non-concatenation uses) in `templates/resolver-mutations.go`

---

## Dependencies & Execution Order

### Phase Dependencies

```text
Setup (P1) ──→ US1 Error Formatting (P3) ──→ US3 Permission Opt (P7) ──→ Polish (P10)
           ├──→ US2 UpdatedAt Fix (P4)    ──→ Polish
           ├──→ US6 OneToMany Clear (P5)  ──→ Polish
           ├──→ US7 OneToOne Unbind (P6)  ──→ Polish
           ├──→ US4 ToOne Conflict (P8)   ──→ Polish
           └──→ US5 Create Simplify (P9)  ──→ Polish
```

- **US1** (P3): Depends on Setup — MVP core, changes error format used by all other phases
- **US2** (P4): Depends on Setup — can run parallel with US1
- **US6** (P5): Depends on Setup — can run parallel with US1/US2
- **US7** (P6): Depends on T003 (verify `InverseRelationship` exists) — can run parallel with US1/US2
- **US3** (P7): Depends on US1 (error format for auth check errors)
- **US4** (P8): Depends on US1 (error format for conflict message)
- **US5** (P9): Independent — can run parallel with everything
- **Polish** (P10): Depends on all USs complete

### Parallel Opportunities

- **T004-T010**: All US1 error formatting tasks target different line ranges
- **US1 + US2 + US6**: Can be implemented in parallel (different template sections)
- **US5**: Independent of all other USs

---

## Implementation Strategy

### MVP First (US1 + US2 + US6)

1. Complete Phase 1: Setup
2. Complete Phase 3-5: US1 + US2 + US6 (all P1)
3. **STOP and VALIDATE**: `go build` passes
4. These three P1 stories fix the most critical issues (data integrity + error quality)

### Incremental Delivery

1. Setup → US1 + US2 + US6 → **Validate** (MVP)
2. + US7 OneToOne Unbind → **Validate**
3. + US3 Permission Optimization → **Validate**
4. + US4 ToOne Conflict → **Validate**
5. + US5 Create Simplification → **Validate**
6. Polish → **Final Validation**

---

## Notes

- All changes are in a single file: `templates/resolver-mutations.go`
- Changes modify the **template string**, not Go logic — edits affect generated output
- Error message format changes are NOT breaking changes (error strings are not API contract)
- US2 (UpdatedAt) and US6 (OneToMany clear) are **data correctness bugs** — highest real-world impact
- US7 (OneToOne unbind) depends on `$rel.InverseRelationship` — T003 must verify this first
- US5 (Create simplification) is the riskiest change — changes zero-value input handling
