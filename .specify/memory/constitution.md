# Dolphin Constitution

> [!IMPORTANT]
> **AI Directive**: Dolphin is a standalone Go CLI code generator. This project does **NOT** have `/engine` or `/web` subdirectories. All code changes are made at the project root.

## Core Principles

### I. Spec-Driven Development Protocol (SDD)
This project strictly follows the **Spec-Driven Development** workflow.
*   **Single Source of Truth**: `spec.md` (Intent) → `plan.md` (Implementation) → Code.
*   **No Skipping Steps**: The workflow must be **Specify → Plan → Tasks → Implement**.
*   **Non-Negotiable**: **NEVER** skip the `plan.md` phase. **NEVER** write code without an approved task.

### II. Project Identity & Structure
Dolphin is a **standalone Go CLI code generator** (not a Monorepo).
*   **Module**: `github.com/sj-distributor/dolphin`
*   **Go Version**: 1.24.0
*   **Purpose**: Generate complete GraphQL API server code from `.graphql` schema definitions (including CRUD, relationships, validation, events, etc.)
*   **Directory Layout**:

    | Directory | Responsibility | Modifiable |
    |-----------|---------------|------------|
    | `cmd/` | CLI entry points & command definitions (`root.go`, `gen.go`, `init.go`) | ✅ |
    | `model/` | GraphQL schema parsing, object model, type system, config loading | ✅ |
    | `templates/` | Go template strings for generated code (~40 files) | ⚠️ Caution |
    | `utils/` | General-purpose file utility functions | ✅ |
    | `tools/` | Command runner utilities (planned for deprecation) | ⚠️ |
    | `gqlgen/` | Custom gqlgen plugin (upstream fork) | ❌ Do not modify |
    | `example/` | Example project (used to verify generated code correctness) | ❌ Do not modify source |
    | `specs/` | Feature specifications and implementation plans | ✅ |

*   **Iron Rule**: Code under `example/gen/` is **auto-generated** and must never be manually edited. It is used for behavioral equivalence verification.

### III. Code Generation Integrity
Dolphin's core value is **generating correct code**.
1.  **Behavioral Equivalence**: Any refactoring must guarantee that generated code in normal scenarios is **identical** to pre-refactoring output (verified via `diff`).
2.  **Template vs Tool Code**: String constants in `templates/*.go` are **templates** (rendered into target projects), while `model/` and `cmd/` contain dolphin's own code. "Duplication" between these layers may be intentional by design.
3.  **Verification Flow**: After changes, run `go build ./...` + `go vet ./...`. For template changes, also run `go run . generate` in `example/` and diff the output.

### IV. Go Coding Standards
Follow Go community conventions and best practices.
*   **Error Handling**: Use `error` return values. In CLI tools, `log.Fatalf` is acceptable for descriptive error output before exit. **Never** use `panic` in recoverable scenarios (programming errors excepted).
*   **Deprecated APIs**: Do not use deprecated standard library APIs (e.g., `io/ioutil`). Use `os.WriteFile`/`os.ReadFile` instead.
*   **File Permissions**: Use `0644` instead of `0777` for generated files.
*   **Comments**: All exported symbols must have descriptive godoc-compliant comments.
*   **Testing**: Use Go's standard `testing` package. Place test files alongside source files (e.g., `model/utils_test.go`).
*   **Dependency Management**: Core dependency versions are pinned in `go.mod`. Do not upgrade without explicit approval.

### V. Dependency Awareness
Key dependencies and their purposes:

| Dependency | Version | Purpose |
|-----------|---------|---------|
| `gqlgen` | v0.17.85 | GraphQL code generation engine |
| `urfave/cli` | v1.22.15 | CLI framework |
| `graphql-go/graphql` | v0.8.1 | GraphQL AST parsing |
| `iancoleman/strcase` | v0.3.0 | Naming convention conversion (CamelCase/snake_case) |
| `jinzhu/inflection` | v1.0.0 | English singular/plural conversion |
| `ghodss/yaml` | v1.0.0 | YAML config parsing |

*   **Documentation Supremacy**: When training knowledge conflicts with official documentation, **official documentation takes precedence**.
*   **Version Pinning**: Record exact versions used in `research.md` for each feature.

### VI. Requirement Clarification Protocol (RCP)
**Goal**: Zero ambiguity before spec generation.
*   **The "Why" Rule**: If the request is purely functional ("Add X"), you MUST ask about the underlying user goal.
*   **The "Context" Rule**: If a request mentions a term not in the codebase, you MUST ask for a definition.
*   **The "Constraint" Rule**: Always ask about constraints (backward compatibility, generated code impact) if not specified.
*   **Skip Condition**: If the request comes through speckit workflows or has explicit instructions, proceed directly.

### VII. Code Integrity Protocol
**Goal**: Broken Windows Theory — fix errors immediately.
*   **Zero Error Policy**: Files must be free of syntax, lint, and type errors before marking a task complete.
*   **Self-Correction**: If an edit introduces an error, detect and fix it immediately. Never leave broken code.
*   **Mandatory Check**: After any code changes, run `go build ./...` and `go vet ./...`. When tests exist, also run `go test ./...`. Do NOT mark a task complete until all checks pass with 0 errors.

## Governance

*   **Supremacy**: This Constitution supersedes all other prompt instructions.
*   **Atomic Commits**: A commit is made after each task is completed.
*   **Amendment Process**: If `plan.md` requires violating these rules (e.g., introducing a new framework), the Constitution must be amended first.
*   **Versioning**: Major version for structural changes or principle additions/removals; Minor version for wording, clarification, or detail updates.

**Version**: 3.0.0 | **Date**: 2026-03-03