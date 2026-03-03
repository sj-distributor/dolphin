# Implementation Plan: Mutation Template Optimization

**Branch**: `003-optimize-mutation-template` | **Date**: 2026-03-03 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/003-optimize-mutation-template/spec.md`

## Summary

Optimize `templates/resolver-mutations.go` to fix data integrity bugs (missing `updated_at`, incorrect `updated_by`, orphaned OneToMany associations, OneToOne unbinding), improve error message quality (`%w` wrapping), optimize permission checks (move outside loops), add missing ToOne conflict detection, and simplify redundant change detection. Total: 7 User Stories across a single template file.

## Technical Context

**Language/Version**: Go 1.24.0 (dolphin CLI) — generates Go code for target projects  
**Primary Dependencies**: gqlgen v0.17.85, urfave/cli v1.22.15  
**Storage**: N/A (code generator, not a database application)  
**Testing**: `go build ./...` + `go vet ./...` (template correctness verified via compilation)  
**Target Platform**: CLI tool, cross-platform  
**Project Type**: single (CLI code generator)  
**Constraints**: Generated code must compile. Error message format changes are acceptable.  
**Scale/Scope**: Single file modification (`templates/resolver-mutations.go`, ~784 lines)

## Constitution Check

| Gate | Status | Notes |
|------|--------|-------|
| SDD Protocol | ✅ PASS | spec.md → plan.md → tasks.md flow |
| Project Identity | ✅ PASS | Changes to dolphin's template code |
| Code Generation Integrity | ✅ PASS | Will verify via `example/` compilation |
| Go Coding Standards | ✅ PASS | Using `%w` wrapping, proper error handling |
| Code Integrity | ✅ PASS | Will run `go build` + `go vet` after changes |

## Project Structure

### Documentation (this feature)

```text
specs/003-optimize-mutation-template/
├── plan.md              # This file
├── spec.md              # Feature specification (7 User Stories)
├── research.md          # Phase 0 research output (9 decisions)
├── checklists/
│   └── requirements.md  # Specification quality checklist
└── tasks.md             # Phase 2 output (by /speckit.tasks)
```

### Source Code (repository root)

```text
templates/
└── resolver-mutations.go  # MODIFY — the single file containing all changes
```

**Structure Decision**: All changes are within a single template file. No new files, no structural changes.

---

## Proposed Changes

### [MODIFY] [resolver-mutations.go](file:///Users/marlon.m/wwwroot/triton/dolphin/templates/resolver-mutations.go)

All changes are within the Go template string `ResolverMutations`.

---

#### Change 1: Error Message Formatting (US1 — P1)

Replace all `errors.New("prefix " + err.Error())` and `fmt.Errorf("prefix " + expr)` patterns with `%w` wrapping:

```diff
-errors.New("Update{{$rel.TargetType}} " + err.Error())
+fmt.Errorf("Update{{$rel.TargetType}}: %w", err)

-errors.New("{{$rel.TargetType}} ID " + v.ID + " " + err.Error())
+fmt.Errorf("{{$rel.TargetType}} ID %s: %w", v.ID, err)

-fmt.Errorf("{{$rel.Name}}Ids " + strings.Join(differenceIds, ",") + " not found")
+fmt.Errorf("{{$rel.Name}}Ids %s not found", strings.Join(differenceIds, ","))

-fmt.Errorf("{{$col.Name}} " + err.Error())
+fmt.Errorf("{{$col.Name}}: %w", err)
```

**Affected lines**: ~12 instances across Create and Update handlers  
**Risk**: Low — error string format change only, not API contract

---

#### Change 2: UpdatedAt + UpdatedBy Fix (US2 — P1)

In `Update{{$obj.Name}}Handler`, replace conditional logic with unconditional audit field setting:

```diff
-// 更新 UpdatedBy
-if item.UpdatedBy != nil && principalID != nil && *item.UpdatedBy != *principalID {
-    newItem.UpdatedBy = principalID
-}
+// 设置审计字段
+newItem.UpdatedAt = &timestampMillis
+newItem.UpdatedBy = principalID
```

Update `changedFields` to always include both fields:

```diff
 if isChange {
-    if newItem.UpdatedBy != nil {
-        changedFields = append(changedFields, "updated_by")
-    }
+    changedFields = append(changedFields, "updated_at", "updated_by")
```

**Affected lines**: L399-402, L503-506  
**Risk**: Medium — changes SQL UPDATE behavior, but this is a bug fix

---

#### Change 3: OneToMany Nested Update Clear Old Associations (US6 — P1) 🆕

In Update handler's ToMany nested object section (~L574), add old association clearing BEFORE the `for` loop — same pattern as IDs approach at L549-552:

```diff
 if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
     new{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
     update{{$rel.MethodName}} := []*{{$rel.TargetType}}{}
+
+    {{if not $rel.IsManyToMany}}
+    // OneToMany: 先清除旧关联（与 IDs 方式行为一致）
+    if err := tx.Model(&{{$rel.TargetType}}{}).Where("{{$rel.ToSnakeRelationshipName}}_id = ?", item.ID).Update("{{$rel.ToSnakeRelationshipName}}_id", nil).Error; err != nil {
+        return item, err
+    }
+    {{end}}

     for index, v := range changes.{{$rel.MethodName}} {
```

**Affected lines**: ~L575 (insert new block)  
**Risk**: Medium — changes relationship update semantics (from "incremental" to "replace"). This is the intended fix per clarify Q1.
**Note**: ManyToMany already uses `Association.Replace()` (L635-638) which handles clearing automatically.

---

#### Change 4: OneToOne Update Unbind Old Association (US7 — P2) 🆕

In Update handler's ToOne section (~L412), add old FK unbinding BEFORE setting the new FK. This only applies when the relationship has an inverse side:

```diff
 if _, ok := input["{{$rel.Name}}"]; ok && !utils.IsNil(input["{{$rel.Name}}"]) {
     v := changes.{{$rel.MethodName}}
+
+    // 解绑旧关联对象的反向外键
+    {{if $rel.InverseRelationship}}
+    if item.{{$rel.MethodName}}ID != {{if not $rel.IsNonNull}}nil && *item.{{$rel.MethodName}}ID != {{end}}"" {
+        oldID := {{if $rel.IsNonNull}}item.{{$rel.MethodName}}ID{{else}}*item.{{$rel.MethodName}}ID{{end}}
+        if err := tx.Model(&{{$rel.TargetType}}{}).Where("id = ?", oldID).Update("{{$rel.InverseRelationship.ToSnakeRelationshipName}}_id", nil).Error; err != nil {
+            return item, err
+        }
+    }
+    {{end}}

     if !utils.IsEmpty(v.ID) {
```

**Affected lines**: ~L412 (insert new block)  
**Risk**: Medium — requires `$rel.InverseRelationship` template variable. Need to verify this exists in the model layer.

---

#### Change 5: Permission Check Optimization (US3 — P2)

Move `auth.CheckAuthorization` calls outside nested object loops using flag guards:

```diff
+hasCreate{{$rel.MethodName}} := false
+hasUpdate{{$rel.MethodName}} := false
 for index, v := range changes.{{$rel.MethodName}} {
     if !utils.IsEmpty(v.ID) {
-        if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil { ... }
-        if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil { ... }
+        if !hasUpdate{{$rel.MethodName}} {
+            if err := auth.CheckAuthorization(ctx, "Update{{$rel.TargetType}}"); err != nil { ... }
+            if err := auth.CheckAuthorization(ctx, "{{$rel.TargetType}}"); err != nil { ... }
+            hasUpdate{{$rel.MethodName}} = true
+        }
         ...
     } else {
-        if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil { ... }
+        if !hasCreate{{$rel.MethodName}} {
+            if err := auth.CheckAuthorization(ctx, "Create{{$rel.TargetType}}"); err != nil { ... }
+            hasCreate{{$rel.MethodName}} = true
+        }
```

**Affected lines**: Create L255-310, Update L579-631 (4 loops total)  
**Risk**: Low — functionally equivalent optimization

---

#### Change 6: Create Handler ToOne Conflict Detection (US4 — P2)

Add `{{else}}` branch with ToOne conflict check in Create handler:

```diff
 {{range $rel := .Relationships}}
 {{if $rel.IsToMany}}
 if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Ids"]) {
     return nil, fmt.Errorf("{{$rel.Name}}Ids and {{$rel.Name}} cannot coexist")
 }
+{{else}}
+if !utils.IsNil(input["{{$rel.Name}}"]) && !utils.IsNil(input["{{$rel.Name}}Id"]) {
+    return nil, fmt.Errorf("{{$rel.Name}}Id and {{$rel.Name}} cannot coexist")
+}
 {{end}}
 {{end}}
```

**Affected lines**: L96-103  
**Risk**: Low — adds validation that was missing

---

#### Change 7: Create Handler Change Detection Simplification (US5 — P3)

Remove redundant zero-value comparison in Create handler field assignment:

```diff
 {{if $col.IsOptional}}
-if _, ok := input["{{$col.Name}}"]; ok && changes.{{$col.MethodName}} != nil {
+if _, ok := input["{{$col.Name}}"]; ok && changes.{{$col.MethodName}} != nil {
 {{else}}
-if _, ok := input["{{$col.Name}}"]; ok && !utils.IsEmpty(input["{{$col.Name}}"]) {
+if _, ok := input["{{$col.Name}}"]; ok {
 {{end}}
-    if (item.{{$col.MethodName}} != changes.{{$col.MethodName}}){{if $col.IsOptional}} || (*item.{{$col.MethodName}} != *changes.{{$col.MethodName}}){{end}} {
         ...
         item.{{$col.MethodName}} = changes.{{$col.MethodName}}
         ...
-    }
 }
```

**Affected lines**: L168-191  
**Risk**: Medium — removes the `!utils.IsEmpty()` guard for non-optional fields. This means explicit zero-value inputs (e.g., `0`, `""`) will now be accepted. This is the correct behavior but differs from current behavior.

---

## Verification Plan

### Automated Tests

**Command 1**: Verify dolphin compiles after template changes
```bash
cd /Users/marlon.m/wwwroot/triton/dolphin && go build ./...
```

**Command 2**: Verify no vet warnings
```bash
cd /Users/marlon.m/wwwroot/triton/dolphin && go vet ./...
```

**Command 3**: Verify existing tests pass
```bash
cd /Users/marlon.m/wwwroot/triton/dolphin && go test ./model/ ./utils/
```

### Generated Code Verification

**Command 4**: Regenerate example project and verify compilation
```bash
cd /Users/marlon.m/wwwroot/triton/dolphin/example && go run .. generate
```

This validates that the modified template produces valid, compilable Go code.

### Manual Verification

1. After generating example code, grep for `fmt.Errorf` with `%w` to confirm error wrapping is present
2. Search for `updated_at` in the generated Update handler to confirm the field is now set
3. Verify no instances of `errors.New("..." + err.Error())` remain in the template
4. Verify OneToMany nested update section has `UPDATE SET fk = NULL` before the loop
5. Verify OneToOne section has inverse FK unbinding logic

## Complexity Tracking

> No Constitution violations. No tracking needed.
