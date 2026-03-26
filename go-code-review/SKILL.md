---
name: "go-code-review"
description: "Professional Go code review skill with context-aware strategies. Invoke when user requests code review, PR review, or quality assessment for Go projects."
---

# Go Code Review Skill

Professional code review skill for Go projects, designed to handle both small changes and large-scale codebases efficiently.

## When to Invoke

- User explicitly requests code review
- User asks for PR/MR review
- User wants quality assessment of Go code
- User mentions "review" in context of Go development
- Before merging significant changes

---

## 0. Review Workflow (Critical)

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

## 1. Quick Reference

### Critical Issues (Must Fix)

| Category | Issues |
|----------|--------|
| **Security** | SQL injection, hardcoded secrets, missing auth |
| **Correctness** | Race conditions, nil pointer dereference, resource leaks |
| **Performance** | N+1 queries, unbounded memory, blocking in hot paths |

### Common Anti-Patterns
- **Panic in business logic**: `panic("something went wrong")`
- **Ignored error**: `_ = file.Close()`
- **Global mutable state**: `var db *sql.DB`
- **Unparameterized SQL**: `db.Query("SELECT * FROM users WHERE id = " + userID)`
- **Context not passed**: `func ProcessData(data []byte) error { ... }`

---

## 2. Project-Specific Focus
- **API Service**: Error handling, input validation, rate limiting
- **CLI Tool**: Flag parsing, error messages, exit codes
- **Library**: API stability, documentation, backward compatibility
- **Microservice**: Service discovery, circuit breakers, observability
- **Data Pipeline**: Memory management, error recovery, idempotency

**Details:** [docs/project-focus.md](docs/project-focus.md)

---

## 3. Incremental Review Strategies

### Strategy A: By Module
Core → API → Data → Tests

### Strategy B: By Priority
Security → Business Logic → Utilities → Tests

### Strategy C: By Dependency Order
Interfaces → Implementations → Consumers → Integration Tests

**For all incremental reviews, use memory.md to track progress.**

---

## 4. Memory-Based Review

**Trigger:** 31+ files or 2000+ lines changed

**Location:** `<project-root>/.review/memory.md`

**Workflow:**
1. Check if memory.md exists → restore or create
2. Determine current phase (Security → Quality → Architecture → Testing)
3. Review batch of 5-10 files
4. Update memory.md with findings
5. Repeat until complete, then generate final report

**Details:** [docs/memory-workflow.md](docs/memory-workflow.md)

---

## 5. Review Output

Use template: [templates/review-report.md](templates/review-report.md)

```markdown
## Code Review Report

### Summary
- Files changed: X
- Risk level: High/Medium/Low
- App Scale: Small/Medium/Large

### Critical Issues (Must Fix)
1. [Security] Description - Location

### Suggestions (Should Fix)
1. [Performance] Description - Location

### Optimization Recommendations
1. [High Priority] Description - Expected ROI: XX%
2. [Medium Priority] Description - Expected ROI: XX%
3. [Low Priority] Description - Expected ROI: XX%

### Over-Engineering Warnings
1. [Over-Engineering] Description - Simplified Alternative: XXX

### Implementation Complexity
1. [Low Complexity] Description - Estimated Effort: X hours
2. [Medium Complexity] Description - Estimated Effort: X hours
3. [High Complexity] Description - Estimated Effort: X hours

### Questions for Author
1. Why was this approach chosen?

### Positive Highlights
- Good error handling in X
```

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
