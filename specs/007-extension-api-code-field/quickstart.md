# Quickstart: Extension API Code Field

## Step 1: Add a custom extension
Modify your GraphQL schema (e.g., `model/test.graphql`) to include a custom extension with the `@entity` title:

```graphql
extend type Query {
  login(input: loginParams!): LoginResult @entity(title: "用户登录")
}
```

## Step 2: Generate Code
Run the dolphin generator:

```bash
make generate
```

## Step 3: Verify api.json
Check `docs/api.json` for the `code` field in the `data` array of the extension entry:

```json
{
  "name": "用户登录",
  "title": "用户登录",
  "api": "login",
  "type": "detail",
  "method": "Query",
  "code": "login(input: loginParams!): LoginResult"
}
```
