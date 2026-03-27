# Data Model: Enum API Metadata

## GraphQL Schema
- **Directives**: `@entity(title: String)` on enum value definitions.
- **Enums**: `Role`, `ClientType`, `Platform` with `@entity` titles on their values.

## Internal structures (api.json typeData)
- **TypeData**:
    - `name`: string (e.g., "ClientType")
    - `fields`: []FieldMetadata
- **FieldMetadata**:
    - `name`: string (e.g., "WEB")
    - `desc`: string (e.g., "网页")
    - `type`: string (e.g., "Enum")
    - `required`: boolean ("false")
    - `validator`: string ("")
