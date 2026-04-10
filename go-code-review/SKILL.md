---
name: go-code-review
description: Professional Go code review skill with context-aware strategies. Invoke when user requests code review, PR review, or quality assessment for Go projects.
---

# Go Code Review Skill

Professional code review skill for Go projects, designed to handle both small changes and large-scale codebases efficiently.

## Table of Contents
- [When to Invoke](#when-to-invoke)
- [Review Workflow](#review-workflow)
- [Quick Reference](#quick-reference)
- [Project-Specific Focus](#project-specific-focus)
- [Incremental Review Strategies](#incremental-review-strategies)
- [Memory-Based Review](#memory-based-review)
- [Review Output](#review-output)
- [Best Practices](#best-practices)
- [Tool Usage](#tool-usage)
- [File Index](#file-index)
- [AskUserQuestion Guidelines](#askuserquestion-guidelines)

## When to Invoke

- User explicitly requests code review
- User asks for PR/MR review
- User wants quality assessment of Go code
- User mentions "review" in context of Go development
- Before merging significant changes

---

## Review Workflow

### Step 1: Scope Assessment
- **Small**: Single file, <200 lines → Direct review
- **Medium**: 200-500 lines, related changes → Focused review
- **Large**: >500 lines, cross-module → Incremental review
- **Project-wide**: Architecture changes → Project analysis first

### Step 2: Context Gathering
- **1-5 files**: Read all changed files
- **6-15 files**: Read primary, summarize secondary
- **16-30 files**: Read changed + direct dependencies
- **31+ files**: Use Memory-Based Review

**Context Rules:**
- Get git diff first (not entire codebase)
- Read ONLY: changed files, direct imports, interface definitions
- Skip: unchanged dependencies, generated code, vendored packages
- **IMPORTANT**: Always check `.gitignore` to exclude non-project files

### Step 3: Execute Review
- **Security Review** → [docs/security-review.md](docs/security-review.md)
- **Quality Review** → [docs/quality-review.md](docs/quality-review.md)
- **Architecture Review** → [docs/architecture-review.md](docs/architecture-review.md)
- **Testing Review** → [docs/testing-review.md](docs/testing-review.md)
- **Scale Assessment** → [docs/scale-assessment.md](docs/scale-assessment.md)
- **Cost-Benefit Analysis** → [docs/cost-benefit-analysis.md](docs/cost-benefit-analysis.md)
- **Over-Engineering Detection** → [docs/over-engineering-detection.md](docs/over-engineering-detection.md)
- **Implementation Complexity Assessment** → [docs/implementation-complexity.md](docs/implementation-complexity.md)
- **Optimization Priority Calibration** → [docs/optimization-priority.md](docs/optimization-priority.md)
- **Automated Analysis** → Use review tool for automated code analysis

### Step 4: Review Tool Integration
- **Location**: `tools/review` (Windows: `tools/review.exe`)
- **Usage**: 
  - `tools/review analyze --path <project-path>`
  - `tools/review report --format markdown`
- **Build**: 
  - `cd review && make install` (构建并安装到tools目录)
  - `cd review && make build` (仅构建到bin目录)
- **Features**: Map preallocation, duplicate code detection, scale detection, markdown reports

---

## Quick Reference

### Critical Issues (Must Fix)

| Category | Issues |
|----------|--------|
| **Security** | SQL injection, hardcoded secrets, missing auth |
| **Correctness** | Race conditions, nil pointer dereference, resource leaks |
| **Performance** | N+1 queries, unbounded memory, blocking in hot paths |

### Common Anti-Patterns
For detailed anti-patterns and examples, see [docs/common-issues.md](docs/common-issues.md)

---

## Project-Specific Focus
- **API Service**: Error handling, input validation, rate limiting, request ID tracing
- **CLI Tool**: Flag parsing, error messages, exit codes
- **Library**: API stability, documentation, backward compatibility
- **Microservice**: Service discovery, circuit breakers, observability
- **Data Pipeline**: Memory management, error recovery, idempotency

**Details:** [docs/project-focus.md](docs/project-focus.md)

---

## Incremental Review Strategies

### Strategy A: By Module
Core → API → Data → Tests

### Strategy B: By Priority
Security → Business Logic → Utilities → Tests

### Strategy C: By Dependency Order
Interfaces → Implementations → Consumers → Integration Tests

**For all incremental reviews, use memory.md to track progress.**

---

## Memory-Based Review

**Trigger:** 31+ files or 2000+ lines changed

**Location:** `<project-root>/.review/`

**Memory Files:**
- **memory-short.md**: Short-term memory for current review cycle
- **memory-long.md**: Long-term memory for cross-cycle knowledge

**Workflow:**
1. Check if memory files exist → restore or create
2. Determine current phase (Security → Quality → Architecture → Testing)
3. Review batch of 5-10 files
4. Update memory files with findings
5. Repeat until complete, then generate final report

**Details:** [docs/memory-workflow.md](docs/memory-workflow.md)

---

## Review Output

### Review Directory Structure
```
<project-root>/.review/
├── memory-short.md      # Short-term memory
├── memory-long.md       # Long-term memory
├── review-rules.md      # Project-specific review rules
├── review-report-*.md   # Review reports
└── archive/             # Archived review files
```

### Report Template
Use template: [templates/review-report.md](templates/review-report.md)

---

## Best Practices

### Before Review
- **Always use askuserquestion** to clarify ambiguous requirements
- Establish review scope and objectives
- Set up review environment with proper memory files

### During Review
- Follow project-specific review rules
- Focus on security, correctness, and performance
- Document both issues and positive aspects
- Maintain consistent terminology

### After Review
- Generate comprehensive review report
- Update memory files with findings
- Provide clear action items and priorities
- Schedule follow-up reviews for critical issues

---

## Tool Usage

For detailed tool usage guidelines, see [docs/tool-usage.md](docs/tool-usage.md)

---

## File Index

```
go-code-review/
├── SKILL.md           # This file - main entry point
├── review/            # Review tool source code
├── docs/              # Review documentation
├── tools/             # Built review tool
└── templates/         # Review templates
```

---

## AskUserQuestion Guidelines

**Use askuserquestion when:**
- Review scope is ambiguous
- Multiple valid approaches exist
- Project-specific rules need clarification
- Security considerations require input
- Performance trade-offs need decision

**Example questions:**
- "What is the intended scope of this review?"
- "Which security aspects should we prioritize?"
- "Are there any specific performance requirements?"
- "What's the expected testing strategy for this change?"

**IMPORTANT**: Always ask questions before making assumptions that could impact the review quality.
