# Feature Specification: Fix Custom Input Directive Processing

**Feature Branch**: `004-fix-input-directive`  
**Created**: 2026-03-27  
**Status**: Draft  
**Input**: User description: "Fix input directive processing for custom Query/Mutation extensions"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Support Custom Query with Input Directives (Priority: P1)

As a developer, I want to define custom Query methods with `@entity` directives on `input` type fields and arguments, so that I can extend the system with custom logic while maintaining metadata consistency.

**Why this priority**: Required for extending the platform with specialized features like login, where custom types and inputs are necessary.

**Independent Test**: Define a custom `Query` in `test.graphql` using an `input` type with `@entity` directives, and run `make generate` successfully.

**Acceptance Scenarios**:

1. **Given** a GraphQL schema with an `input` type containing `@entity` directives, **When** running `make generate`, **Then** the process should complete without `gqlgen` errors.
2. **Given** an `extend type Query` with a field having arguments with directives, **When** running `make generate`, **Then** the final schema should be valid.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST strip `@entity`, `@relationship`, `@column`, and `@validator` directives from `InputObjectDefinition` nodes in the schema printer.
- **FR-002**: System MUST strip dolphin-specific directives from `InputValueDefinition` (arguments) of fields in both `ObjectDefinition` and `TypeExtensionDefinition`.
- **FR-003**: System MUST NOT crash when encountering unknown directives on user-defined input types.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `make generate` in the `example` project completes successfully after adding the `login` example.
- **SC-002**: The generated `schema.graphqls` contains clean definitions without unresolved internal directives.
