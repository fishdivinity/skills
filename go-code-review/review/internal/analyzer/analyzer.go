package analyzer

// AnalysisResult 表示分析结果
type AnalysisResult struct {
	Issues    []Issue
	Scale     string
	FilePath  string
	FileCount int
}

// Issue 表示一个优化机会
type Issue struct {
	Type        string
	Description string
	FilePath    string
	Line        int
	Severity    string
}

// Analyzer 分析器接口
type Analyzer interface {
	Analyze(path string, diff bool, scale string) (AnalysisResult, error)
}

// analyzer 分析器实现
type analyzer struct{}

// NewAnalyzer 创建一个新的分析器
func NewAnalyzer() Analyzer {
	return &analyzer{}
}

// Analyze 执行代码分析
func (a *analyzer) Analyze(path string, diff bool, scale string) (AnalysisResult, error) {
	var issues []Issue
	var files []string
	var err error

	// 根据diff参数决定分析所有文件还是只分析更改的文件
	if diff {
		// 只分析更改的Go文件
		changedFiles, err := getChangedGoFiles(path)
		if err != nil {
			// 如果git diff失败，回退到分析所有文件
			files, err = findGoFiles(path)
		} else {
			// 构建完整的文件路径
			for _, file := range changedFiles {
				fullPath := path + "/" + file
				files = append(files, fullPath)
			}
		}
	} else {
		// 分析所有Go文件
		files, err = findGoFiles(path)
	}

	if err != nil {
		return AnalysisResult{}, err
	}

	// 对每个文件执行分析
	for _, file := range files {
		// 检测map预分配
		mapIssues, err := detectMapPreallocation(file)
		if err == nil {
			issues = append(issues, mapIssues...)
		}

		// 检测重复代码
		duplicateIssues, err := detectDuplicateCode(file)
		if err == nil {
			issues = append(issues, duplicateIssues...)
		}
	}

	// 如果没有指定规模，自动检测
	if scale == "" {
		scale = detectScale(len(files), len(issues))
	}

	return AnalysisResult{
		Issues:    issues,
		Scale:     scale,
		FilePath:  path,
		FileCount: len(files),
	}, nil
}

// detectScale 检测应用程序规模
func detectScale(fileCount int, _ int) string {
	if fileCount < 10 {
		return "small"
	} else if fileCount < 50 {
		return "medium"
	} else {
		return "large"
	}
}
