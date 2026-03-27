# Implementation Plan: Extension API Metadata

**Branch**: `005-extension-api-metadata` | **Date**: 2026-03-27 | **Spec**: [specs/005-extension-api-metadata/spec.md](file:///Users/marlon.m/wwwroot/dolphin/specs/005-extension-api-metadata/spec.md)
**Input**: Feature specification from `/specs/005-extension-api-metadata/spec.md`

## Summary

Dolphin currently generates basic metadata for `extend type Query` and `extend type Mutation` in `api.json`, but the `typeData` field is empty. This feature will update the `model` package to traverse the GraphQL AST for these extensions, extract input argument types and return type structures, and expose them to the `templates/graphql.go` template for documentation generation.

## Technical Context

**Language/Version**: Go 1.25.5
**Primary Dependencies**: gqlgen v0.17.85, graphql-go/graphql v0.8.1
**Storage**: N/A
**Testing**: go test ./model/...
**Target Platform**: CLI
**Project Type**: single
**Performance Goals**: N/A
**Constraints**: Must maintain backward compatibility for existing CRUD documentation.
**Scale/Scope**: Impacts all custom GraphQL extensions in `api.json`.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] Standalone Go CLI code generator? Yes.
- [x] No /engine or /web subdirectories? Yes.
- [x] Spec-Driven Development followed? Yes.

## Project Structure

### Documentation (this feature)

```text
specs/005-extension-api-metadata/
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
├── model.go             # Handle document-level extensions
├── model.object.go      # Update Object structure to hold extension metadata
├── model.object-field.go # Update ObjectField to include input/output types
templates/
└── graphql.go           # Render extracted metadata into api.json
```

**Structure Decision**: Standard single project structure.

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| None | | |
