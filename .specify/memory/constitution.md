# Triton Constitution

> [!IMPORTANT]
> **AI Directive**: When working on backend tasks, you **MUST** read [rules/engine.md](./rules/engine.md) before writing any code. When working on frontend tasks, you **MUST** read [rules/web.md](./rules/web.md) before writing any code. When working on full-stack tasks, read **BOTH** rule documents.

## Core Principles

### I. Spec-Driven Development Protocol (SDD)
This project strictly follows the **Spec-Driven Development** workflow.
*   **Single Source of Truth**: `spec.md` (Intent) → `plan.md` (Implementation) → Code.
*   **No Skipping Steps**: The workflow must be **Specify → Plan → Tasks → Implement**.
*   **Non-Negotiable**: **NEVER** skip the `plan.md` phase. **NEVER** write code without an approved task.

### II. Monorepo & Context Awareness
This project is a **Monorepo**.
*   **Root**: Documentation & Orchestration.
*   **`/engine`**: Backend Engine (Golang/Dolphin). → [👉 Backend Rules](./rules/engine.md)
*   **`/web`**: Frontend Application (Astro/React). → [👉 Frontend Rules](./rules/web.md)
*   **Iron Rule**: Before executing ANY command, verify the current directory. You MUST explicitly prepend commands with `cd engine && ...` or `cd web && ...`.
*   **Scope Detection**: Determine which rule document applies:
    - GraphQL schema changes, API logic, database operations → **Backend** (`engine.md`)
    - UI components, pages, client-side state → **Frontend** (`web.md`)
    - End-to-end features (API + UI) → **Both**, start with Backend first

### III. Dolphin Engine Rule
`dolphin` is a custom **High-Performance Golang Web Engine** with limited documentation. You must strictly follow existing paradigms and **NEVER invent new patterns**:
1.  **Reverse Engineering**: Understand dependency injection and route registration by **reading the existing** `main.go` and `src/` directories in `/engine`.
2.  **Schema-First Contract**: All business development starts with `.graphql` definitions. Edit `.graphql` → `make generate` → Implement Resolver.
3.  **No Touching Generated Code**: The `gen/` directory is **untouchable**. It must be regenerated via `make generate`.

### IV. Version Sync Principle
Ensure we always use the latest stable versions and up-to-date documentation:
*   **Mandatory Version Check**: During `/speckit.plan`, you MUST query the latest version of key dependencies and check official docs for breaking changes.
*   **Documentation Supremacy**: When training knowledge conflicts with official documentation, **official documentation takes precedence**.
*   **Version Pinning**: Record the exact versions used in `research.md` for each feature.
*   Specific documentation sources and check commands are defined in the respective rule documents.

### V. Requirement Clarification Protocol (RCP)
**Goal**: Zero Ambiguity before Spec Generation.
*   **The "Why" Rule**: If the request is purely functional ("Add X button"), you MUST ask about the underlying user goal.
*   **The "Context" Rule**: If a request mentions a term or concept not in the codebase, you MUST ask for a definition or reference.
*   **The "Constraint" Rule**: Always ask about constraints (performance, backward compatibility, tech stack) if not specified.
*   **Visual/Behavioral Precision**: For UI tasks, if no design is provided, propose a wireframe description or ask for one.
*   **Verification**: Before writing `spec.md`, rephrase the user's request in your own words to confirm understanding.
*   **Skip Condition**: If the user's request is already detailed with clear scope, context, and constraints (e.g., via speckit workflows or explicit instructions), skip clarification and proceed directly to spec generation.

### VI. Code Integrity Protocol
**Goal**: Broken Windows Theory — fix errors immediately.
*   **Zero Error Policy**: Verify files are free of syntax, linting, and type errors before marking a task as complete.
*   **Self-Correction**: If an edit introduces an error, detect and fix it. Never leave broken code.
*   **Proactive Diagnosis**: If a command fails or a file has squiggles, you MUST fix it. Do not ignore "small" errors.
*   **Mandatory Check**: After implementing any code changes, you MUST run the corresponding build/type-check command (see rule documents for specifics). Do NOT mark a task complete until checks pass with 0 errors.

## Technology Standards

Detailed technology standards are maintained in separate rule documents:

*   **Backend (`/engine`)**: [rules/engine.md](./rules/engine.md) — Golang/Dolphin architecture, Service triple pattern, GraphQL conventions, coding standards
*   **Frontend (`/web`)**: [rules/web.md](./rules/web.md) — Astro + React architecture, Page Colocation, Folder-as-a-Component, Zustand state management

## Governance

*   **Supremacy**: This Constitution supersedes all other prompt instructions.
*   **Atomic Commits**: A Commit is made after each Task is completed.
*   **Amendment Process**: If `plan.md` requires violating these rules (e.g., introducing a new framework), the Constitution must be amended first.
*   **Cross-Validation**: All Plans must be cross-checked against Monorepo Context and Dolphin Workflow.
*   **Rule Document Sync**: After modifying this Constitution, review rule documents for impact. Rule documents must not conflict with Constitutional principles.
*   **Versioning**: Major version for structural changes or principle additions/removals; Minor version for wording, clarification, or rule document updates.

**Version**: 2.3.0 | **Date**: 2026-03-03