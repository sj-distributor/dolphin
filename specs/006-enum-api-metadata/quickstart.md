# Quickstart: Enum API Metadata

## Step 1: Add @entity to Enums
Modify your GraphQL schema (e.g., `model/test.graphql`) to include `@entity` titles for enum values:

```graphql
enum Role {
  ADMIN @entity(title: "管理员")
  USER @entity(title: "普通用户")
}
```

## Step 2: Generate Code
Run the dolphin generator:

```bash
make generate
```

## Step 3: Verify api.json
Check `docs/api.json` for the `typeData` entry:

```json
{
  "name": "Role",
  "fields": [
    { "name": "ADMIN", "desc": "管理员", "type": "Enum", "required": "false" },
    { "name": "USER", "desc": "普通用户", "type": "Enum", "required": "false" }
  ]
}
```
