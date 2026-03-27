# Quickstart: Fix Custom Input Directive Processing

## Verification Steps
1. Navigate to the `example` directory.
2. Ensure `model/test.graphql` contains a custom `input` and `extend type Query` extension.
3. Run `make generate`.
4. Verify that the command completes without errors.
5. Check `gen/schema.graphqls` to ensure dolphin-specific directives are removed.
