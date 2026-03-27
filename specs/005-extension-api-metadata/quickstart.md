# Quickstart: Extension API Metadata

## Verification Steps
1. Navigate to the `example` directory.
2. Run `make generate`.
3. Open `example/docs/api.json`.
4. Locate the `login` endpoint in the extensions section.
5. Verify that `typeData` contains entries for `loginParams` and `LoginResult`.
6. Confirm that the `fields` within `typeData` correctly list the GraphQL fields and their metadata (e.g., descriptions from `@entity`).
