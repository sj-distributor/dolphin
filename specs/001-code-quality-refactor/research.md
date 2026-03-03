# Research: 代码质量重构

**Feature**: 001-code-quality-refactor  
**Date**: 2026-03-03

## R1: isExist() 函数 Bug

**Decision**: `os.IsExist(err)` 与 `!os.IsNotExist(err)` 语义不同，当前实现有 Bug  
**Rationale**: `os.IsExist` 检查"文件已存在"错误（如创建已存在文件时），**不是** `os.IsNotExist` 的反向。正确判断文件是否存在应使用 `err == nil` 或 `!os.IsNotExist(err)`。当前 `isExist()` 在文件存在且无错误时返回 `false`（因为 `os.IsExist(nil)` 返回 false）  
**Alternatives considered**: 使用 `errors.Is` — 过于复杂；直接用 `os.Stat(path); err == nil` 最简洁

## R2: panic vs error 返回策略

**Decision**: 将所有非测试代码中的 `panic` 替换为 `error` 返回  
**Rationale**: Go 惯例是使用 error 返回值处理可恢复错误。panic 应仅用于真正不可恢复的编程错误。dolphin 中的 panic 都用于"对象未找到""关系缺失"等运行时可恢复场景  
**Alternatives considered**: 保留 panic 并添加 recover — 增加复杂性且不符合 Go 惯例

## R3: ioutil 废弃替代

**Decision**: `ioutil.WriteFile` → `os.WriteFile`，`ioutil.ReadFile` → `os.ReadFile`  
**Rationale**: Go 1.16 起 `io/ioutil` 包已废弃，功能已迁移至 `os` 和 `io` 包。项目使用 Go 1.24，完全支持  
**Alternatives considered**: 无，标准库直接替换

## R4: 重复函数统一策略

**Decision**: 将公共工具函数统一到 `utils` 包或 `model` 包中  
**Rationale**: 当前 `RunInteractiveInDir`、`GetRandomString`、`IndexOf` 各有两份实现。`templates/` 中的是生成到目标项目的模板代码，`model/` 和 `tools/` 中的是 dolphin 自身使用的代码，需分别处理  
**Alternatives considered**: 保持现状加注释 — 仍有维护不一致风险

## R5: 模板字符串中的代码 vs dolphin 自身代码

**Decision**: 区分"dolphin 自身代码"和"模板生成的目标代码"两个层面  
**Rationale**: `templates/` 目录下的 Go 字符串常量是**模板**，会被渲染为目标项目的代码。其中的"重复"（如 `IndexOf`、`GetRandomString`）实际是生成代码中需要的，不应简单删除。但 dolphin 自身的 `model/utils.go` 与 `tools/run.go` 中的重复应消除  
**Alternatives considered**: 全部统一 — 不可行，模板代码和 dolphin 代码运行在不同项目中

## R6: 单元测试框架选择

**Decision**: 使用 Go 标准 `testing` 包 + `go test`  
**Rationale**: 项目当前无任何测试。Go 标准库已足够，无需引入额外测试框架  
**Alternatives considered**: testify — 增加依赖，且项目规模不需要

## R7: Go 版本与工具链

**Version**: Go 1.24.0（来自 `go.mod`）  
**gqlgen**: v0.17.85  
**urfave/cli**: v1.22.15  
**Decision**: 不升级依赖版本，仅重构代码质量
