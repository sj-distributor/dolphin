# Feature Specification: Enum API Metadata

**Feature Branch**: `006-enum-api-metadata`  
**Created**: 2026-03-27  
**Status**: Draft  
**Input**: User description: "Implement @entity support for enums and display titles in api.json"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - See Enum Value Titles in Documentation (Priority: P1)

As a developer using the generated API, I want to see the human-readable titles for each enum value in the generated `api.json` file, so that the frontend can display these titles instead of technical keys.

**Why this priority**: Enums are widely used for status, types, and categories. Displaying technical keys (like `ADMIN`) instead of user-friendly names (like `管理员`) makes the documentation less useful for frontend integration.

**Independent Test**: Add `@entity(title: "...")` to enum values in `test.graphql`, run `make generate`, and verify that `api.json`'s `typeData` includes these titles for the enum values.

**Acceptance Scenarios**:

1. **Given** an enum with `@entity(title: "...")` on its values, **When** running `make generate`, **Then** the `api.json` should contain the titles for each enum value in the corresponding `typeData` entry.
2. **Given** an enum value without a title, **When** running `make generate`, **Then** the `api.json` should fallback to the value's name.

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST support the `@entity` directive on enum value definitions in the GraphQL schema.
- **FR-002**: System MUST extract the `title` argument from the `@entity` directive on enum values.
- **FR-003**: System MUST include enum value metadata (name and title) in the `typeData` field of `api.json`.
- **FR-004**: System MUST ensure that `GetTypeData()` in `ObjectField` also extracts metadata for enums related to the field.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `example/docs/api.json` contains titles for `ClientType` and `Platform` enum values.
- **SC-002**: The `typeData` entry for an enum has a `fields` list where each field represents an enum value and contains its `desc` (title).
