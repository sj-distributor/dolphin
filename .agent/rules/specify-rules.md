---
trigger: always_on
---

# triton Development Guidelines

Auto-generated from all feature plans. Last updated: 2026-01-15

## 语言偏好
- 所有交互和文档使用简体中文

## Active Technologies
- TypeScript (via Astro) + Astro, React, Tailwind CSS, shadcn/ui utilities (002-upgrade-web-deps)
- TypeScript 5.x, Astro 5.x, React 19.x. (001-scalable-web-arch)
- LocalStorage (for client-side persistence if needed). (001-scalable-web-arch)
- TypeScript 5.x (via Astro) + Astro 5.x, React 19.x, React Router 7.12.0, TailwindCSS 4.1.18, shadcn/ui (001-feature-based-web)
- LocalStorage (客户端主题偏好) (001-feature-based-web)
- TypeScript 5.x + Zustand 5.0.10, Immer 10.x (002-state-theme-fix)
- LocalStorage (主题偏好持久化) (002-state-theme-fix)
- TypeScript 5.x + i18next, react-i18next, i18next-browser-languagedetector (003-i18n-centralized)
- LocalStorage (语言偏好持久化) (003-i18n-centralized)
- TypeScript 5.x (tsx 运行) + 无新增依赖（使用 Node.js 原生 fetch） (004-ai-translation-generator)
- TypeScript 5.x + Astro 5.x, React 19.x (010-web-arch-refactor)
- TypeScript 5.x (via Astro) + Astro 5.x, React 19.x, Tailwind CSS 4.x, shadcn/ui (011-storefront-homepage)
- 静态 JSON 文件 (`web/src/data/`) (011-storefront-homepage)
- TypeScript 5.x (via Astro) + Astro 5.x, React 19.x, Apollo Client 4.x, Zustand 5.x, react-hook-form 7.71.2, zod 4.3.6, @hookform/resolvers 5.2.2 (016-user-registration)
- LocalStorage (JWT 持久化 via Zustand persist) (016-user-registration)
- TypeScript 5.x + Astro 5.x, React 19.x, `react-router` 7.x, `@tanstack/react-table` (8.21.3) (017-admin-scaffold)
- Zustand (客户端状态) (017-admin-scaffold)
- N/A (Documentation Only) (001-folder-component-arch)
- TypeScript 5.x (via Astro) + Astro 5.x, React 19.x, @tanstack/react-table, react-hook-form, @apollo/clien (001-folder-component-arch)
- Go 1.25.5 + dolphin (code generation), gqlgen, golang-jwt/v5, **go-redis/redis/v9** (NEW) (018-user-login)
- MySQL (via GORM/dolphin DAL) + **Redis** (Active Token store) (018-user-login)
- Go 1.25.5 + golang-jwt/v5, gqlgen (existing, no new deps) (019-auth-refactor)
- N/A (no storage changes) (019-auth-refactor)
- Go 1.24.0 (dolphin CLI) — generates Go code for target projects + gqlgen v0.17.85, urfave/cli v1.22.15 (003-optimize-mutation-template)
- N/A (code generator, not a database application) (003-optimize-mutation-template)

- TypeScript 5.x (via Astro defaults) + Astro 5.x, React 18.x, Tailwind CSS 3.x, shadcn/ui utilities (001-astro-frontend-setup)

## Project Structure

```text
backend/
frontend/
tests/
```

## Commands

npm test && npm run lint

## Code Style

TypeScript 5.x (via Astro defaults): Follow standard conventions

## Recent Changes
- 003-optimize-mutation-template: Added Go 1.24.0 (dolphin CLI) — generates Go code for target projects + gqlgen v0.17.85, urfave/cli v1.22.15
- 003-optimize-mutation-template: Added Go 1.24.0 (dolphin CLI) — generates Go code for target projects + gqlgen v0.17.85, urfave/cli v1.22.15
- 019-auth-refactor: Added Go 1.25.5 + golang-jwt/v5, gqlgen (existing, no new deps)


<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
