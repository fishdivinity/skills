# Memory-Based Review Workflow

**For large projects with >30 files or >2000 lines changed**

---

## Why Memory File?

| Problem | Solution |
|---------|----------|
| Context overflow | Persist review state to file |
| Cannot resume | Restore progress from memory |
| Duplicate work | Track reviewed files |
| Lost findings | Accumulate issues across rounds |

---

## Memory File Location

```
<project-root>/.review/memory.md
```

Use template: `templates/memory.md`

---

## Memory File Structure

```markdown
# Review Memory

## Project Info
- Type: API Service / CLI / Library / Microservice
- Total Files: X
- Lines Changed: +X / -Y
- Review Started: YYYY-MM-DD

## Progress
| Phase | Status | Files Reviewed | Issues Found |
|-------|--------|----------------|--------------|
| Security | Done | 5 | 3 |
| Quality | In Progress | 3 | 2 |
| Architecture | Pending | 0 | 0 |
| Testing | Pending | 0 | 0 |

## Reviewed Files
- [x] cmd/server/main.go
- [x] internal/handlers/auth.go
- [ ] internal/service/user.go
- [ ] internal/repository/order.go

## Issues Found

### Critical (Must Fix)
| ID | Category | File | Line | Description | Status |
|----|----------|------|------|-------------|--------|
| C1 | Security | auth.go | 45 | SQL injection | Fixed |
| C2 | Security | config.go | 12 | Hardcoded secret | Pending |

### Suggestions (Should Fix)
| ID | Category | File | Line | Description | Status |
|----|----------|------|------|-------------|--------|
| S1 | Quality | user.go | 89 | Missing error wrap | Pending |
| S2 | Architecture | order.go | 120 | Layer violation | Pending |

## Context Notes
- Uses chi router (not gorilla/mux)
- Custom error types in internal/errors/
- PostgreSQL with pgx driver

## Next Steps
1. Continue Quality review: internal/service/
2. Start Architecture review
3. Review test coverage
```

---

## Complete Workflow

```
┌─────────────────────────────────────────────────────────────┐
│                    START REVIEW                             │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Check if .review/memory.md exists?                         │
│  - YES: Read and restore context                            │
│  - NO: Create new memory file                               │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Determine current phase based on memory:                   │
│  1. Security Review                                         │
│  2. Quality Review                                          │
│  3. Architecture Review                                     │
│  4. Testing Review                                          │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Review next batch of files (5-10 files per round)          │
│  - Read files not yet reviewed                              │
│  - Apply relevant checklist                                 │
│  - Record findings                                          │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Update memory.md:                                          │
│  - Mark reviewed files                                      │
│  - Add new issues                                           │
│  - Update progress                                          │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  All phases complete?                                       │
│  - NO: Return to next phase (can pause here)                │
│  - YES: Generate final report                               │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  FINAL OUTPUT:                                              │
│  - Generate review-report.md                                │
│  - Optionally archive or delete memory.md                   │
└─────────────────────────────────────────────────────────────┘
```

---

## Batch Size Guidelines

| Project Size | Files per Round | Estimated Rounds |
|--------------|-----------------|------------------|
| 10-30 files | 5-8 | 2-4 |
| 30-60 files | 8-10 | 4-8 |
| 60-100 files | 10-15 | 8-12 |
| 100+ files | 15-20 | 10+ |

---

## Reading Memory Before Each Session

1. **Project Info** - Understand project type and scope
2. **Progress Overview** - Know which phase to continue
3. **Reviewed Files** - Skip already reviewed files
4. **Issues Found** - Don't duplicate findings
5. **Context Notes** - Remember patterns and technologies
6. **Next Steps** - Know exactly where to continue

---

## Updating Memory After Each Batch

1. **Progress Overview**
   - Increment Files Reviewed count
   - Increment Issues Found count
   - Update Status if phase changed

2. **File Tracking**
   - Mark reviewed files: `- [ ]` → `- [x]`

3. **Issues Found**
   - Add new issues with unique IDs (C1, C2, H1, S1, etc.)
   - Include file, line, description

4. **Context Notes**
   - Add new patterns observed
   - Note technologies discovered

5. **Review History**
   - Add entry with date, phase, files, issues

6. **Next Steps**
   - Update to reflect remaining work

---

## Memory File Best Practices

### DO
- Update memory after EVERY batch
- Use consistent issue IDs (C1, C2, H1, S1, Q1)
- Keep Context Notes concise but useful
- Archive completed reviews with timestamp

### DON'T
- Don't skip updating memory between batches
- Don't delete memory until review is complete
- Don't duplicate issue IDs
- Don't leave status as "In Progress" when done

---

## Example Memory Update Sequence

**Before Batch 1:**
```markdown
## Progress Overview
| Phase | Status | Files Reviewed | Issues Found |
|-------|--------|----------------|--------------|
| Security | Pending | 0 | 0 |
```

**After Batch 1 (reviewed 5 files, found 2 issues):**
```markdown
## Progress Overview
| Phase | Status | Files Reviewed | Issues Found |
|-------|--------|----------------|--------------|
| Security | In Progress | 5 | 2 |

## Issues Found
| ID | Category | File | Line | Description | Status |
|----|----------|------|------|-------------|--------|
| C1 | Security | auth.go | 45 | SQL injection in login | Pending |
| C2 | Security | config.go | 12 | Hardcoded API key | Pending |
```

**After Batch 2 (reviewed 3 more files, found 1 issue):**
```markdown
## Progress Overview
| Phase | Status | Files Reviewed | Issues Found |
|-------|--------|----------------|--------------|
| Security | Done | 8 | 3 |

## Issues Found
| ID | Category | File | Line | Description | Status |
|----|----------|------|------|-------------|--------|
| C1 | Security | auth.go | 45 | SQL injection in login | Pending |
| C2 | Security | config.go | 12 | Hardcoded API key | Pending |
| C3 | Security | handler.go | 89 | Missing auth check | Pending |
```

---

## Memory File Commands

```bash
# Check if memory exists
ls .review/memory.md

# View current progress
head -50 .review/memory.md

# Count remaining files
grep -c "^\- \[ \]" .review/memory.md

# Archive completed review
mv .review/memory.md .review/memory-$(date +%Y%m%d).md
```

---

## When to Use Memory File

| Condition | Use Memory |
|-----------|------------|
| Files changed > 30 | Yes |
| Lines changed > 2000 | Yes |
| Cross-module changes | Yes |
| Multi-session review needed | Yes |
| Architecture-level review | Yes |
