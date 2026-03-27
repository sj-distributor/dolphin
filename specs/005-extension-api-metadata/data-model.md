# Data Model: Extension API Metadata

This feature updates Dolphin's internal model to support richer metadata for custom extensions.

## Internal Model Changes
- `Model`: Added `GetDefinition(name string) ast.Node` for type discovery.
- `ObjectField`:
  - Added `GetTypeData() []TypeData` for metadata extraction.
  - `TypeData` struct: Contains `Name` and associated `Fields`.
  - `FieldMetadata` struct: Contains `Name`, `Desc`, `Type`, and `Required`.
