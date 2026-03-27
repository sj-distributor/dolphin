# Data Model: Fix Custom Input Directive Processing

This feature focuses on the internal code generation logic of Dolphin and does not introduce new entities to the target project.

## Generator Changes
- Refactor `model/printer.go` to recursively strip directives from ALL GraphQL AST nodes.
