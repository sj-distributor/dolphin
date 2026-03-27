# Research: Extension API Code Field

## Decision: Signature Generation
- **Approach**: Add a `Signature()` method to `model.ObjectField`.
- **Logic**: Use `graphql-go/graphql/language/printer.Print` to render the `ast.FieldDefinition`.
- **Directive Handling**: Use the existing `StripDirectives` logic (or similar) from `model/printer.go` to ensure Dolphin-specific directives are not included in the output.

## Rationale
- Leveraging the standard GraphQL printer ensures consistent and valid syntax for lists, non-nulls, and complex arguments.
- Adding the method to `ObjectField` makes it reusable for other documentation parts if needed.

## Findings: templates/graphql.go
- The `data` array in `api.json` is generated in `templates/graphql.go`.
- I'll need to update the JSON template to include the `"code": "{{ .Signature }}"` field.

## Next Steps
- Implement `Signature()` in `model/model.object-field.go`.
- Update `templates/graphql.go` to include the `code` field.
