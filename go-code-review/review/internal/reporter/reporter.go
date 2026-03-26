package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"review/internal/analyzer"
)

// Reporter 报告生成器接口
type Reporter interface {
	Generate(format string, input string) (string, error)
	GenerateFromResult(format string, result analyzer.AnalysisResult) (string, error)
}

// reporter 报告生成器实现
type reporter struct{}

// NewReporter 创建一个新的报告生成器
func NewReporter() Reporter {
	return &reporter{}
}

// Generate 生成报告
func (r *reporter) Generate(format string, input string) (string, error) {
	// 从文件读取分析结果
	var result analyzer.AnalysisResult

	if input != "" {
		// 读取文件内容
		content, err := os.ReadFile(input)
		if err != nil {
			return "", fmt.Errorf("error reading input file: %v", err)
		}

		// 解析JSON格式的分析结果
		if err := json.Unmarshal(content, &result); err != nil {
			return "", fmt.Errorf("error parsing input file: %v", err)
		}
	}

	switch format {
	case "markdown":
		return generateMarkdownReport(result), nil
	default:
		return generateTextReport(result), nil
	}
}

// GenerateFromResult 从分析结果生成报告
func (r *reporter) GenerateFromResult(format string, result analyzer.AnalysisResult) (string, error) {
	switch format {
	case "markdown":
		return generateMarkdownReport(result), nil
	default:
		return generateTextReport(result), nil
	}
}

// generateMarkdownReport 生成markdown格式报告
func generateMarkdownReport(result analyzer.AnalysisResult) string {
	report := `# Code Review Report

## Summary
`

	report += fmt.Sprintf("- **Files Analyzed**: %d\n", result.FileCount)
	report += fmt.Sprintf("- **Issues Found**: %d\n", len(result.Issues))
	report += fmt.Sprintf("- **Scale**: %s\n", result.Scale)

	report += "\n## Optimization Opportunities\n"

	if len(result.Issues) == 0 {
		report += "\nNo issues found.\n"
	} else {
		for _, issue := range result.Issues {
			report += fmt.Sprintf("\n### %s\n", issue.Type)
			report += fmt.Sprintf("- **Description**: %s\n", issue.Description)
			report += fmt.Sprintf("- **File**: %s\n", issue.FilePath)
			report += fmt.Sprintf("- **Line**: %d\n", issue.Line)
			report += fmt.Sprintf("- **Severity**: %s\n", issue.Severity)
		}
	}

	report += "\n## Recommendations\n"

	// 根据规模生成不同的建议
	switch result.Scale {
	case "small":
		report += "\n- Focus on code quality and readability\n"
		report += "- Simplify monitoring to essential metrics only\n"
		report += "- Avoid complex caching strategies\n"
		report += "- Prioritize readability over micro-optimizations\n"
	case "medium":
		report += "\n- Balance performance and readability\n"
		report += "- Implement standard caching with basic monitoring\n"
		report += "- Optimize critical paths only\n"
	case "large":
		report += "\n- Prioritize performance and scalability\n"
		report += "- Implement comprehensive monitoring\n"
		report += "- Optimize all performance-critical paths\n"
	}

	report += "\n- Regularly review code for optimization opportunities\n"
	report += "- Follow Go best practices\n"

	return report
}

// generateTextReport 生成文本格式报告
func generateTextReport(result analyzer.AnalysisResult) string {
	report := `Code Review Report
==================

Summary:
`

	report += fmt.Sprintf("- Files Analyzed: %d\n", result.FileCount)
	report += fmt.Sprintf("- Issues Found: %d\n", len(result.Issues))
	report += fmt.Sprintf("- Scale: %s\n", result.Scale)

	report += "\nOptimization Opportunities:\n"

	if len(result.Issues) == 0 {
		report += "\nNo issues found.\n"
	} else {
		for _, issue := range result.Issues {
			report += fmt.Sprintf("\n%s:\n", issue.Type)
			report += fmt.Sprintf("  Description: %s\n", issue.Description)
			report += fmt.Sprintf("  File: %s\n", issue.FilePath)
			report += fmt.Sprintf("  Line: %d\n", issue.Line)
			report += fmt.Sprintf("  Severity: %s\n", issue.Severity)
		}
	}

	report += "\nRecommendations:\n"

	// 根据规模生成不同的建议
	switch result.Scale {
	case "small":
		report += "\n- Focus on code quality and readability\n"
		report += "- Simplify monitoring to essential metrics only\n"
		report += "- Avoid complex caching strategies\n"
		report += "- Prioritize readability over micro-optimizations\n"
	case "medium":
		report += "\n- Balance performance and readability\n"
		report += "- Implement standard caching with basic monitoring\n"
		report += "- Optimize critical paths only\n"
	case "large":
		report += "\n- Prioritize performance and scalability\n"
		report += "- Implement comprehensive monitoring\n"
		report += "- Optimize all performance-critical paths\n"
	}

	report += "\n- Regularly review code for optimization opportunities\n"
	report += "- Follow Go best practices\n"

	return report
}
