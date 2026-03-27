# Research: Fix Custom Input Directive Processing

## Decision: Refactor `PrintSchema` in `model/printer.go`

### Rationale
- `model/definition.query.go` and `model/definition.mutation.go` are only responsible for generating standard CRUD fields for entities. They do not handle custom `extend type Query/Mutation` definitions.
- Custom extensions are preserved in the `ast.Document` and printed to the final schema via `model/printer.go`.
- `gqlgen` fails because the printed schema contains unknown dolphin-specific directives on `input` types and arguments that were never stripped.
- Fixing it in `PrintSchema` ensures that *all* definitions (including manual extensions and inputs) are cleaned before being passed to `gqlgen`.

### Alternatives Considered

#### 1. Modify `definition.query.go` and `definition.mutation.go` to absorb extensions
- **Rejected because**: This would require dolphin to manually merge extensions into the main `Query`/`Mutation` types, replicate `gqlgen`'s merging logic, and handle complex AST transformations. It's more complex and prone to bugs than simply cleaning the AST before printing.

#### 2. Define dolphin directives in the final schema
- **Rejected because**: We don't want dolphin's internal metadata directives (like `@entity(title: "...")`) to leak into the production GraphQL API. They are for code generation time only.

## Key Findings
- `InputObjectDefinition` and `InputValueDefinition` nodes are currently ignored by the stripping logic in `PrintSchema`.
- Custom `extend type Query` fields are processed for their field-level directives, but their arguments (`InputValueDefinition`) are not.
- Adding a recursive AST cleaning function or a more comprehensive loop in `PrintSchema` will solve both issues.
