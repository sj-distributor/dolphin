# Research: Enum Metadata Extraction

## Decision: Metadata Structure for Enums
- **Approach**: Treat Enums as a special "Type" in `api.json`'s `typeData`.
- **Mapping**: Each Enum value will be represented as a "Field" in the `Fields` slice of `TypeData`.
    - `Name`: Enum value name (e.g., `ADMIN`)
    - `Desc`: Title from `@entity(title: "...")`
    - `Type`: "Enum" (or leave empty as it's not a type)
    - `Required`: "false" (not applicable)

## Rationale
- Reusing the `TypeData` and `FieldMetadata` structs avoids introducing new JSON structures for the documentation, making it easier for the frontend to consume.

## Alternatives Considered
- Creating a separate `EnumData` structure.
- **Rejected**: Adds complexity to templates and frontend parsers.

## Findings: GraphQL AST
- `ast.EnumDefinition` contains `Values []*ast.EnumValueDefinition`.
- `ast.EnumValueDefinition` has `Directives []*ast.Directive`.
- We can use the same `getValidatorFromDirectives` or create a `getEnumTitle` helper in `model/model.object-field.go`.

## Next Steps
- Implement `Case *ast.EnumDefinition` in `GetTypeData`.
- Update `GetDefinition` in `model/model.go` to ensure all enum definitions are reachable (already implemented, but verify for custom ones).
