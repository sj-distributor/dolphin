# Implementation Plan: Fix Custom Input Directive Processing

**Branch**: `004-fix-input-directive` | **Date**: 2026-03-27 | **Spec**: [specs/004-fix-input-directive/spec.md](file:///Users/marlon.m/wwwroot/dolphin/specs/004-fix-input-directive/spec.md)
**Input**: Feature specification from `/specs/004-fix-input-directive/spec.md`

## Summary

The current implementation of `model/printer.go` fails to strip dolphin-specific directives (like `@entity`) from `InputObjectDefinition` (input types) and `InputValueDefinition` (field arguments). This leads to invalid GraphQL schemas being passed to `gqlgen`, causing generation failures. The proposed fix is to refactor `PrintSchema` to recursively clean all AST nodes in the document.

## Technical Context

**Language/Version**: Go 1.25.5
**Primary Dependencies**: gqlgen v0.17.85, graphql-go/graphql v0.8.1
**Storage**: N/A
**Testing**: go test ./model/...
**Target Platform**: CLI
**Project Type**: single
**Performance Goals**: N/A
**Constraints**: Must not break existing CRUD generation or other extensions.
**Scale/Scope**: Impacts all GraphQL schema generation.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] Standalone Go CLI code generator? Yes.
- [x] No /engine or /web subdirectories? Yes.
- [x] Spec-Driven Development followed? Yes.
- [x] Go Version 1.24.0+? Yes.

## Project Structure

### Documentation (this feature)

```text
specs/004-fix-input-directive/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── tasks.md             # Phase 2 output
```

### Source Code (repository root)

```text
model/
├── printer.go           # Primary target for directive stripping logic
├── definition.query.go  # Check for potential improvements in query generation
└── definition.mutation.go # Check for potential improvements in mutation generation
```

**Structure Decision**: Standard single project structure.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None | | |
