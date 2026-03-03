# Feature Specification: Mutation Template Optimization

**Feature Branch**: `003-optimize-mutation-template`  
**Created**: 2026-03-03  
**Status**: Draft  
**Input**: Deep analysis of `templates/resolver-mutations.go` identified 9+ issues across the Create and Update mutation template. Fix all: code duplication, permission check optimization, error formatting, UpdatedBy/UpdatedAt logic, change detection redundancy, and relationship handling correctness.

## Clarifications

### Session 2026-03-03

- Q: OneToMany 嵌套对象更新时是否需要清除旧关联？ → A: Option A — 先清除旧关联再设置新关联，与 IDs 方式行为保持一致
- Q: OneToOne 更新时是否需要解绑旧关联对象的反向外键？ → A: Option A — 更新前先将旧关联对象的反向外键设为 nil
- Q: OneToMany 嵌套更新中外键双写是否需要优化？ → A: Option B — 保持现状，两次写操作确保外键正确，性能影响可忽略

## User Scenarios & Testing

### User Story 1 - Error Message Quality Improvement (Priority: P1)

As a developer debugging generated API code, I want error messages to use Go's `%w` error wrapping convention so that I can use `errors.Is`/`errors.As` to programmatically inspect error chains instead of parsing raw strings.

**Why this priority**: Error handling quality directly affects developer experience and production debugging. This is the safest change with the highest immediate value.

**Independent Test**: Generate code, trigger an error (e.g., invalid relationship ID), verify the error message uses `fmt.Errorf("...: %w", err)` format and can be unwrapped with `errors.Is`.

**Acceptance Scenarios**:

1. **Given** a Create mutation with an invalid relationship ID, **When** the handler returns an error, **Then** the error message uses `fmt.Errorf` with `%w` verb for proper error wrapping
2. **Given** an Update mutation where a nested object update fails, **When** the error propagates, **Then** callers can use `errors.Is` to inspect the underlying cause
3. **Given** a ToMany relationship with missing IDs, **When** validation fails, **Then** the error message uses `fmt.Errorf` with structured format instead of string concatenation

---

### User Story 2 - UpdatedAt/UpdatedBy Logic Fix (Priority: P1)

As a data consumer, I want the `updated_at` and `updated_by` fields to always reflect the latest modification so that audit trails are accurate and complete.

**Why this priority**: Missing or incorrect audit fields is a data integrity issue that affects all downstream consumers (reports, sync, compliance).

**Independent Test**: Perform an Update mutation, query the entity, verify `updated_at` and `updated_by` reflect the current operation.

**Acceptance Scenarios**:

1. **Given** an entity with `updated_by = null` (never updated before), **When** a different user performs an update, **Then** `updated_by` is set to the current user's principal ID
2. **Given** an entity updated by user A, **When** user A updates the same entity again, **Then** `updated_at` is refreshed to the current timestamp (even though `updated_by` stays the same)
3. **Given** an Update mutation that changes at least one field, **When** the mutation completes, **Then** both `updated_at` and `updated_by` are included in the SQL UPDATE statement

---

### User Story 3 - Permission Check Optimization (Priority: P2)

As a system operator, I want permission checks to execute at most once per relationship type per mutation so that unnecessary repeated authorization calls are eliminated, improving response time for mutations with many related objects.

**Why this priority**: Performance optimization that becomes significant when creating/updating entities with many nested relationships (e.g., an order with 50 line items).

**Independent Test**: Create an entity with 10 nested ToMany objects, verify only one pair of `auth.CheckAuthorization` calls is made for that relationship type (not 10 pairs).

**Acceptance Scenarios**:

1. **Given** a Create mutation with N nested objects for a single relationship, **When** the handler processes the nested objects, **Then** `auth.CheckAuthorization("Create{{Type}}")` and `auth.CheckAuthorization("{{Type}}")` are called at most once (before the loop), not N times
2. **Given** an Update mutation with mixed create/update nested objects, **When** processing, **Then** Update and Create permission checks each execute once, not per-item
3. **Given** a user without "Create" permission for a related type, **When** they try to create nested objects, **Then** the error is returned immediately before any database operations

---

### User Story 4 - Error Handling Consistency with ToOne Conflict Detection (Priority: P2)

As a developer, I want the Update mutation to validate ToOne relationship field conflicts (e.g., passing both `user` object and `userId` simultaneously) so that ambiguous inputs are caught early with clear error messages.

**Why this priority**: The Create handler only validates ToMany conflicts. The Update handler validates both ToMany and ToOne conflicts, but the Create handler is missing ToOne conflict validation for input consistency.

**Independent Test**: On Create mutation, pass both `user` (nested object) and `userId` (ID field) for the same ToOne relationship, verify an error is returned.

**Acceptance Scenarios**:

1. **Given** a Create mutation where both `{{rel.Name}}` and `{{rel.Name}}Id` are provided for a ToOne relationship, **When** the handler validates input, **Then** a clear error is returned: "{{rel.Name}}Id and {{rel.Name}} cannot coexist"
2. **Given** an Update mutation with the same conflict, **When** the handler validates input, **Then** the same error format is returned (existing behavior preserved)

---

### User Story 5 - Create Handler Change Detection Simplification (Priority: P3)

As a code maintainer, I want the Create handler to skip unnecessary change detection comparison logic (comparing against a zero-value struct) so that generated code is cleaner and easier to understand.

**Why this priority**: Code readability improvement. The comparison is semantically meaningless in Create (zero-value vs input), though functionally harmless.

**Independent Test**: Generate code, compare Create handler output after optimization with original — verify field assignment still works correctly.

**Acceptance Scenarios**:

1. **Given** a Create mutation with normal field input, **When** the handler assigns fields, **Then** all provided fields are set on the entity without the redundant zero-value comparison
2. **Given** a Create mutation where a field value is explicitly set to the Go zero value (e.g., `0` for int, `""` for string), **When** processed, **Then** the zero value IS assigned (fixing a subtle edge case where the old comparison would skip zero-value assignments)

---

### User Story 6 - OneToMany Nested Update Old Association Cleanup (Priority: P1)

As a data consumer, I want OneToMany nested object updates to clear old associations that are not in the submitted list so that orphaned foreign key references do not remain, keeping data consistent with the IDs-based update behavior.

**Why this priority**: Data integrity bug — submitting a partial nested list leaves "ghost" associations (old records still pointing to the parent entity). This breaks the implicit "replace" semantic that users expect.

**Independent Test**: Entity has children [A, B, C]. Update mutation sends nested objects [A, D(new)]. After mutation, B and C's foreign key should be null (no longer pointing to parent).

**Acceptance Scenarios**:

1. **Given** an entity with OneToMany children [A, B, C], **When** Update mutation sends nested objects [A], **Then** B and C's foreign keys are set to null
2. **Given** an entity with OneToMany children [A, B], **When** Update mutation sends nested objects [A, B, C(new)], **Then** A and B are updated, C is created with correct FK, no orphans exist
3. **Given** an entity with ManyToMany associations, **When** Update mutation sends nested objects, **Then** `Association.Replace()` handles cleanup automatically (no change needed)

---

### User Story 7 - OneToOne Update Unbind Old Association (Priority: P2)

As a data consumer, I want OneToOne relationship updates to unbind the old associated object before binding the new one so that strict one-to-one mapping is maintained and no "one-to-many" data inconsistency occurs.

**Why this priority**: If the old OneToOne associated object retains its reverse foreign key while the parent now points to a new object, the old object still "thinks" it's associated — violating OneToOne semantics.

**Independent Test**: User has Profile_A. Update changes to Profile_B. After update, Profile_A's reverse FK (if any) should be null.

**Acceptance Scenarios**:

1. **Given** an entity with OneToOne relationship to Object_A, **When** Update mutation changes to Object_B (existing), **Then** Object_A's reverse foreign key is set to null before Object_B's FK is updated
2. **Given** an entity with OneToOne relationship to Object_A, **When** Update mutation creates a new Object_C (no ID), **Then** Object_A's reverse foreign key is set to null before Object_C is created and linked

---

### Edge Cases

- What happens when a Create mutation contains both nested objects AND IDs for the same ToMany relationship? (Already handled: returns conflict error)
- What happens when `UpdatedAt` auto-update is handled by GORM tags vs manual setting? (Verified: template manually manages timestamps, GORM auto-update NOT used)
- What happens when permission check fails mid-way through a batch of nested objects? (Transaction rollback handles this)
- How does OneToMany ID reassignment behave under concurrent updates? (Existing transaction isolation handles this)
- What happens when OneToMany nested update sends an empty list? (Should clear all associations — same as IDs approach with empty array)
- What happens when OneToOne update target is the same object? (Should detect no change and skip unbind/rebind)

## Requirements

### Functional Requirements

- **FR-001**: All `errors.New("prefix " + err.Error())` and `fmt.Errorf("prefix " + expr)` patterns MUST be replaced with `fmt.Errorf("prefix: %w", err)` or `fmt.Errorf("prefix %s", expr)` as appropriate
- **FR-002**: Update handler MUST always set `updated_at` to the current timestamp when any field changes
- **FR-003**: Update handler MUST always set `updated_by` to the current principal ID, regardless of whether the previous `updated_by` was the same user or nil
- **FR-004**: Permission checks for nested relationship processing MUST execute once per relationship type per mutation, not once per nested object
- **FR-005**: Create handler MUST validate ToOne relationship field conflicts (same pattern as Update handler)
- **FR-006**: Create handler field assignment SHOULD remove the zero-value comparison guard where it provides no benefit
- **FR-007**: All changes MUST preserve behavioral equivalence for existing valid inputs (no breaking changes to API contract)
- **FR-008**: Update handler's OneToMany nested object processing MUST clear old associations not present in the submitted list before processing new/updated objects (consistent with IDs approach)
- **FR-009**: Update handler's OneToOne processing MUST unbind the old associated object's reverse foreign key before binding the new one
- **FR-010**: OneToMany nested update FK double-write pattern (recursive Update handler call + separate FK update) MUST be preserved as-is for safety

### Assumptions

- GORM does NOT auto-update `updated_at` via tags — the template is responsible for setting it explicitly (confirmed via code analysis)
- `updated_at` uses millisecond Unix timestamp format (consistent with `created_at`)
- The authorization service `auth.CheckAuthorization` is idempotent and safe to call once per type
- Template changes affect ALL entities equally (they share the same template)
- OneToMany "clear old associations" uses `UPDATE SET fk = NULL` pattern (not DELETE), preserving child records

## Success Criteria

### Measurable Outcomes

- **SC-001**: Zero instances of `errors.New("..." + err.Error())` or `fmt.Errorf("..." + expr + "...")` string concatenation in the mutation template
- **SC-002**: Update handler sets `updated_at` and `updated_by` on every successful mutation that modifies at least one field
- **SC-003**: For a mutation with N nested objects, permission checks execute O(1) times per relationship type, not O(N)
- **SC-004**: `go build ./...` and `go vet ./...` pass with zero errors after changes
- **SC-005**: Generated code compiles successfully in the `example/` project
- **SC-006**: OneToMany nested object update clears associations not in the submitted list (no orphaned foreign keys)
- **SC-007**: OneToOne update unbinds old associated object before binding new one
