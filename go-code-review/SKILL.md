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

| Scope | Indicators | Strategy |
|-------|------------|----------|
| **Small** | Single file, <200 lines | Direct review |
| **Medium** | 200-500 lines, related changes | Focused review |
| **Large** | >500 lines, cross-module | **Incremental review** |
| **Project-wide** | Architecture changes | **Project analysis first** |

### Step 2: Context Gathering

| File Count | Strategy |
|------------|----------|
| 1-5 files | Read all changed files |
| 6-15 files | Read primary, summarize secondary |
| 16-30 files | Read changed + direct dependencies |
| 31+ files | **Use Memory-Based Review** |

**Context Rules:**
- Get git diff first (not entire codebase)
- Read ONLY: changed files, direct imports, interface definitions
- Skip: unchanged dependencies, generated code, vendored packages

### Step 3: Execute Review

1. **Security Review** → [docs/security-review.md](docs/security-review.md)
2. **Quality Review** → [docs/quality-review.md](docs/quality-review.md)
3. **Architecture Review** → [docs/architecture-review.md](docs/architecture-review.md)
4. **Testing Review** → [docs/testing-review.md](docs/testing-review.md)
5. **Scale Assessment** → [docs/scale-assessment.md](docs/scale-assessment.md)
6. **Cost-Benefit Analysis** → [docs/cost-benefit-analysis.md](docs/cost-benefit-analysis.md)
7. **Over-Engineering Detection** → [docs/over-engineering-detection.md](docs/over-engineering-detection.md)
8. **Implementation Complexity Assessment** → [docs/implementation-complexity.md](docs/implementation-complexity.md)

---

## 1. Quick Reference

### Critical Issues (Must Fix)

| Category | Issues |
|----------|--------|
| **Security** | SQL injection, hardcoded secrets, missing auth |
| **Correctness** | Race conditions, nil pointer dereference, resource leaks |
| **Performance** | N+1 queries, unbounded memory, blocking in hot paths |

### Common Anti-Patterns

```go
// WRONG: panic in business logic
panic("something went wrong")

// WRONG: ignored error
_ = file.Close()

// WRONG: global mutable state
var db *sql.DB

// WRONG: unparameterized SQL
db.Query("SELECT * FROM users WHERE id = " + userID)

// WRONG: context not passed
func ProcessData(data []byte) error { ... }
```

---

## 2. Project-Specific Focus

| Project Type | Focus Areas |
|--------------|-------------|
| **API Service** | Error handling, input validation, rate limiting |
| **CLI Tool** | Flag parsing, error messages, exit codes |
| **Library** | API stability, documentation, backward compatibility |
| **Microservice** | Service discovery, circuit breakers, observability |
| **Data Pipeline** | Memory management, error recovery, idempotency |

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
├── SKILL.md                       # This file - main entry point
├── docs/
│   ├── security-review.md         # Security review checklist
│   ├── quality-review.md          # Code quality checklist
│   ├── architecture-review.md     # Architecture patterns review
│   ├── testing-review.md          # Testing requirements
│   ├── common-issues.md           # Common Go anti-patterns
│   ├── project-focus.md           # Project-specific review focus
│   ├── inconsistency-detection.md # Pattern inconsistency detection
│   ├── memory-workflow.md         # Memory-based review workflow
│   ├── scale-assessment.md        # Application scale assessment
│   ├── cost-benefit-analysis.md   # Cost-benefit analysis
│   ├── over-engineering-detection.md # Over-engineering detection
│   └── implementation-complexity.md # Implementation complexity assessment
└── templates/
    ├── review-report.md           # Review report template
    └── memory.md                  # Memory file template
```
