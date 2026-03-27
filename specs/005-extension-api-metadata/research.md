# Research: Extension API Metadata

## Problem Analysis
- Custom GraphQL extensions (`extend type Query/Mutation`) are documented in `api.json`, but their nested type structures (input types and return types) are currently missing.
- The `model` package primarily focuses on entities (marked with `@entity`) but needs to support extracting metadata from all types used in custom extensions.
- `InputObjectDefinition` and skipped `ObjectDefinition` nodes contain the metadata needed for `typeData`.

## Decision: Add Generic Type Lookup and Metadata Extraction
- **Model Lookup**: Extend `Model` with a `GetDefinition(name string) ast.Node` method to find any GraphQL definition in the AST document.
- **Field Metadata Extraction**: Add `GetTypeData()` to `ObjectField`. This method will recursively (or iteratively) collect definitions for types used as arguments or return values.
- **Unified Metadata Struct**: Create `TypeData` and `FieldMetadata` structs to provide a clean interface for templates.

## Implementation Details
1. **Model Updates**:
   - `model/model.go`: Implement `GetDefinition`.
   - `model/model.object-field.go`: Implement `GetTypeData` and helper structs.
2. **Template Updates**:
   - `templates/graphql.go`: Iterate over `GetTypeData()` to populate `typeData` in `api.json`.

## Alternatives Considered
- **Mapping all types at load time**: Rejected because it adds overhead for many types that might not be used in documentation. Lookup by name on demand is sufficient.
- **Using gqlgen's internal model**: Rejected as Dolphin is a generator that feeds *into* gqlgen; it should remain independent and use the parsed AST.
