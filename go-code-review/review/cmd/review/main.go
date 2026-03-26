package main

import (
	"flag"
	"fmt"
	"os"
	"review/internal/analyzer"
	"review/internal/reporter"
)

func main() {
	// 定义命令行标志
	analyzeCmd := flag.NewFlagSet("analyze", flag.ExitOnError)
	analyzePath := analyzeCmd.String("path", ".", "Path to analyze")
	analyzeDiff := analyzeCmd.Bool("diff", false, "Analyze git diff")
	analyzeScale := analyzeCmd.String("scale", "", "Override detected scale (small/medium/large)")

	reportCmd := flag.NewFlagSet("report", flag.ExitOnError)
	reportFormat := reportCmd.String("format", "markdown", "Report format")
	reportInput := reportCmd.String("input", "", "Input file with analysis results")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "analyze":
		analyzeCmd.Parse(os.Args[2:])
		handleAnalyzeCommand(*analyzePath, *analyzeDiff, *analyzeScale)
	case "report":
		reportCmd.Parse(os.Args[2:])
		handleReportCommand(*reportFormat, *reportInput)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: review [command] [flags]")
	fmt.Println("Commands:")
	fmt.Println("  analyze: Analyze code for optimization opportunities")
	fmt.Println("  report: Generate report from analysis")
	fmt.Println("  help: Show this help message")
	fmt.Println()
	fmt.Println("Analyze command flags:")
	fmt.Println("  -path string: Path to analyze (default \"\")")
	fmt.Println("  -diff: Analyze git diff")
	fmt.Println("  -scale string: Override detected scale (small/medium/large)")
	fmt.Println()
	fmt.Println("Report command flags:")
	fmt.Println("  -format string: Report format (default \"markdown\")")
	fmt.Println("  -input string: Input file with analysis results")
}

func handleAnalyzeCommand(path string, diff bool, scale string) {
	fmt.Printf("Analyzing path: %s\n", path)
	fmt.Printf("Analyze git diff: %v\n", diff)
	fmt.Printf("Override scale: %s\n", scale)

	// 创建分析器
	analyzer := analyzer.NewAnalyzer()

	// 执行分析
	results, err := analyzer.Analyze(path, diff, scale)
	if err != nil {
		fmt.Printf("Error analyzing code: %v\n", err)
		os.Exit(1)
	}

	// 打印分析结果
	fmt.Println("Analysis completed successfully!")
	fmt.Printf("Analyzed %d files\n", results.FileCount)
	fmt.Printf("Found %d optimization opportunities\n", len(results.Issues))
	if len(results.Issues) == 0 {
		fmt.Println("No optimization opportunities found")
	}
	fmt.Printf("Detected application scale: %s\n", results.Scale)

	// 生成报告
	reporter := reporter.NewReporter()
	report, err := reporter.GenerateFromResult("markdown", results)
	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		os.Exit(1)
	}

	// 输出报告
	fmt.Println("\nReport:")
	fmt.Println(report)
}

func handleReportCommand(format string, input string) {
	fmt.Printf("Generating report in %s format\n", format)
	fmt.Printf("Input file: %s\n", input)

	// 创建报告生成器
	reporter := reporter.NewReporter()

	// 生成报告
	report, err := reporter.Generate(format, input)
	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		os.Exit(1)
	}

	// 输出报告
	fmt.Println("Report generated successfully!")
	fmt.Println(report)
}
