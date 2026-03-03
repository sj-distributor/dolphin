# Research: Mutation Template Optimization

**Feature**: 003-optimize-mutation-template  
**Date**: 2026-03-03 (Updated after clarify session)

## R1: UpdatedAt handling in mutation template

**Decision**: Template manually manages timestamps — GORM auto-update is NOT used  
**Rationale**: Searched for `UpdatedAt` in `templates/resolver-mutations.go`. Found 4 occurrences, ALL on **related objects** (`v.UpdatedAt = &timestampMillis`), NONE on the main entity in UpdateHandler. The Create handler sets `item.CreatedAt = timestampMillis` (L92), confirming manual management. The `updated_at` field is a Unix millisecond timestamp (`int64`), not `time.Time`, so GORM's `autoUpdateTime` doesn't apply.  
**Conclusion**: US2 is confirmed as a real bug — `updated_at` is never set on the main entity during Update.

## R2: UpdatedBy logic analysis

**Decision**: Current logic at L400-402 only sets `UpdatedBy` when `item.UpdatedBy != nil && *item.UpdatedBy != *principalID`  
**Rationale**: This means: (1) if `UpdatedBy` was never set (nil), it stays nil even after update; (2) if the same user updates again, `UpdatedBy` isn't refreshed. Both are incorrect — `UpdatedBy` should always reflect the last modifier.  
**Fix**: Unconditionally set `newItem.UpdatedBy = principalID`

## R3: Error formatting patterns in template

**Decision**: Replace string concatenation with `fmt.Errorf` + `%w`  
**Rationale**: Found 12+ instances of `errors.New("prefix " + err.Error())` pattern. Go 1.13+ supports `%w` for error wrapping, enabling `errors.Is`/`errors.As`. The generated code uses Go modules, so `%w` is safe.  
**Alternatives considered**: Keep `errors.New` — loses error chain information

## R4: Permission check placement

**Decision**: Move `auth.CheckAuthorization` calls before the nested object loop  
**Rationale**: Permission is checked per-type, not per-object. In the current template, a loop over N nested objects calls `CheckAuthorization` N times with the same arguments. Moving before the loop is functionally equivalent and reduces N calls to 1.  
**Alternatives considered**: Caching in auth service — more complex, not dolphin's responsibility

## R5: Create handler ToOne conflict detection

**Decision**: Add ToOne conflict detection to Create handler (matching Update handler)  
**Rationale**: Update handler L387-391 validates `{{rel.Name}}` vs `{{rel.Name}}Id` conflict for ToOne relationships. Create handler only validates ToMany conflicts (L96-103). This is an inconsistency — the same ambiguous input should be caught in both Create and Update.

## R6: Template change scope and risk

**Decision**: All changes are in `templates/resolver-mutations.go` only (template string modifications)  
**Rationale**: Changes affect the generated Go code, not dolphin's own logic. Verification requires regenerating `example/` and confirming compilation. Behavioral equivalence only matters for the "no-error path" — error message format changes are acceptable breaking changes in the error string (not in API contract).

## R7: OneToMany nested update old association clearing (NEW from clarify)

**Decision**: Add `UPDATE SET fk = NULL` for old associations before processing nested objects  
**Rationale**: User confirmed Option A — OneToMany nested object update should clear old associations not in the submitted list, matching the IDs approach (L549-552). Currently, the nested object approach (L574-642) only processes submitted objects without cleaning up orphans.  
**Implementation**: Before the `for` loop, execute `tx.Model(&{{TargetType}}{}).Where("xxx_id = ?", item.ID).Update("xxx_id", nil)` — same pattern as IDs approach at L550.  
**Alternatives considered**: Option B (keep as "incremental update" semantic) — rejected for data consistency; Option C (add `replace` parameter) — over-engineering.

## R8: OneToOne update unbind old association (NEW from clarify)

**Decision**: Unbind old associated object's reverse FK before binding new one  
**Rationale**: User confirmed Option A — when changing a OneToOne relationship from Object_A to Object_B, the template should set Object_A's reverse FK to nil before proceeding. This prevents "one-to-many" data inconsistency where both the old and new objects believe they're associated.  
**Implementation**: Before updating the FK, query the current FK value, then `UPDATE SET reverse_fk = NULL WHERE id = old_fk_value`. Only applies when the relationship has an inverse side.  
**Alternatives considered**: Option B (rely on business layer) — too fragile; Option C (only manual management) — inconsistent with auto-generated code philosophy.

## R9: OneToMany nested update FK double-write pattern (NEW from clarify)

**Decision**: Keep current dual-write pattern (recursive Update handler + separate FK update)  
**Rationale**: User confirmed Option B. The pattern at L596-603 is safe: (1) `r.Handlers.Update{{TargetType}}()` updates field values via `changedFields`; (2) separate `tx.Model(v).Update("xxx_id", item.ID)` ensures FK is set. These are independent SQL statements in the same transaction. Consolidating them risks missing the FK update.  
**Alternatives considered**: Option A (inject FK into input for unified handler processing) — risky since Update handler might not include FK in `changedFields`.
