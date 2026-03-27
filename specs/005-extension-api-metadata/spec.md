# Feature Specification: Extension API Metadata

**Feature Branch**: `005-extension-api-metadata`  
**Created**: 2026-03-27  
**Status**: Draft  
**Input**: User description: "Implement input and result field extraction for custom extend type Query/Mutation to include in api.json"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - See Custom API Metadata in Documentation (Priority: P1)

As a developer using the generated API, I want to see the input parameters and return type details for custom GraphQL extensions (extend type Query/Mutation) in the generated `api.json` file, so that I can automatically generate documentation or client code for these custom endpoints.

**Why this priority**: High priority because custom extensions are currently "black boxes" in the generated documentation, lacking crucial parameter and return type metadata.

**Independent Test**: Add an `extend type Query` with custom `input` and `result` types to the GraphQL schema, run `make generate`, and verify that `example/docs/api.json` contains the populated `typeData` for that extension.

**Acceptance Scenarios**:

1. **Given** a custom query extension with multiple input arguments, **When** running `make generate`, **Then** the generated `api.json` should list all arguments in the `typeData` field for that API.
2. **Given** a custom query extension with a complex return type, **When** running `make generate`, **Then** the `api.json` should provide information about the return type structure.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST extract input argument metadata (name, type, required status) from `extend type Query` and `extend type Mutation` fields in the GraphQL AST.
- **FR-002**: System MUST store this extracted metadata in the `ObjectField` or a similar model structure.
- **FR-003**: System MUST provide the extracted metadata to the `templates/graphql.go` template engine.
- **FR-004**: System MUST populate the `typeData` field in the generated `api.json` for all custom extensions.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `example/docs/api.json` contains non-empty `typeData` for the `login` extension added in `test.graphql`.
- **SC-002**: The `typeData` correctly reflects the names and types of the input fields of the custom `input` type used in the extension.
