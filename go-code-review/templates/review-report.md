# Code Review Report Template

## Summary

| Metric | Value |
|--------|-------|
| Files Changed | X |
| Lines Added | X |
| Lines Removed | X |
| Risk Level | High / Medium / Low |
| Review Status | Approved / Changes Requested / Needs Discussion |

---

## Overview

Brief description of what this change accomplishes.

---

## Critical Issues (Must Fix)

### 1. [Category] Issue Title

**Location:** `file.go:123`

**Description:**
Detailed description of the issue.

**Current Code:**
```go
// problematic code
```

**Suggested Fix:**
```go
// improved code
```

**Impact:** Why this needs to be fixed before merge.

---

## Suggestions (Should Fix)

### 1. [Category] Suggestion Title

**Location:** `file.go:456`

**Description:**
Description of the improvement opportunity.

**Current Code:**
```go
// current code
```

**Suggested Improvement:**
```go
// suggested code
```

**Rationale:** Why this improvement would be beneficial.

---

## Questions for Author

1. **Question 1?**
   - Context for the question
   - What needs clarification

2. **Question 2?**
   - Context for the question
   - What needs clarification

---

## Positive Highlights

- Good error handling in [specific area]
- Well-structured tests for [specific functionality]
- Clean separation of concerns in [specific module]
- Good documentation for [specific feature]

---

## Checklist Results

### Security Review

- [x] No SQL injection vulnerabilities
- [x] No hardcoded secrets
- [x] Authentication properly implemented
- [ ] Rate limiting needs review

### Code Quality Review

- [x] All errors handled
- [x] No race conditions
- [x] Resources properly cleaned up
- [ ] Some functions too long

### Architecture Review

- [x] Proper layer separation
- [x] Dependency injection used
- [ ] Middleware inconsistency detected

### Testing Review

- [x] Unit tests present
- [x] Edge cases covered
- [ ] Integration tests missing

---

## Scale-Based Recommendations

### Application Scale Assessment
- **Scale Category**: Small / Medium / Large
- **Justification**: [Brief explanation based on user count, data volume, etc.]

### Priority Calibration

#### High Priority
- [High Priority] Description - Expected ROI: XX% - Effort: X hours

#### Medium Priority
- [Medium Priority] Description - Expected ROI: XX% - Effort: X hours

#### Low Priority
- [Low Priority] Description - Expected ROI: XX% - Effort: X hours

### Over-Engineering Warnings
- [Over-Engineering] Description - Simplified Alternative: XXX - Reason: Not appropriate for {App Scale} application

### Implementation Guidelines
- **For {App Scale} application**: Focus on {key focus areas}
- **Estimated Implementation Time**: X-X hours (calibrated for {App Scale} application)

---

## Next Steps

1. Address critical issues
2. Respond to questions
3. Consider suggestions
4. Re-request review when ready

---

## Reviewer Notes

Additional context or observations that may be helpful.
