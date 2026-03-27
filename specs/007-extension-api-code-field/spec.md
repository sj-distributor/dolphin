# Feature Specification: Extension API Code Field

**Feature Branch**: `007-extension-api-code-field`  
**Created**: 2026-03-27  
**Status**: Draft  
**Input**: User description: "自定义扩展的接口生成出来的json缺少了data里面的code字段...直接把login(input: loginParams!): LoginResult放入到code字段里面就快可以。"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - See Signature Code in API Doc (Priority: P1)

As an API consumer, I want to see the exact GraphQL signature of a custom extension in the `api.json` document so that I can easily understand how to call it without hunting through schema files.

**Why this priority**: Essential for developers using the generated API documentation to integrate with custom extensions.

**Independent Test**: Generate `api.json` and verify that entries in the `data` array for custom extensions (Query/Mutation) contain a `code` field with the correct GraphQL signature.

**Acceptance Scenarios**:

1. **Given** a GraphQL schema with `extend type Query { login(input: loginParams!): LoginResult }`, **When** I run the code generator, **Then** the `api.json` entry for `login` should have `"code": "login(input: loginParams!): LoginResult"`.
2. **Given** a custom mutation extension, **When** I run the generator, **Then** the `code` field should reflect the mutation's signature.

### Edge Cases

- What happens if the extension has no arguments? (The `code` field should still show the name and return type).
- How are complex types (lists, non-nulls) displayed in the `code` field? (They should be rendered as standard GraphQL syntax).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST extract the GraphQL signature (name, arguments, and return type) for entries in the `data` array of `api.json` that correspond to custom extensions.
- **FR-002**: System MUST include this signature in the `code` field of those `data` objects.
- **FR-003**: The signature MUST NOT include Dolphin-specific directives (e.g., `@entity`, `@validator`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: All custom Query/Mutation extensions in `api.json` have a non-empty `code` field containing their GraphQL signature.
- **SC-002**: The `code` field accurately matches the GraphQL signature defined in the source `.graphql` files.
