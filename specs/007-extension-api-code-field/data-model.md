# Data Model: Extension API Code Field

## API JSON structure
The `api.json` file's `data` array entries for extensions will include a `code` field:

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

## Model Changes
- `model.ObjectField`: Add `Signature() string` method.
- `templates/graphql.go`: Update the template to include the `code` field.
