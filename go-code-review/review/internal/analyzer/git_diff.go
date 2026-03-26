package analyzer

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// getGitDiff 获取git diff输出
func getGitDiff(path string) (string, error) {
	// 验证path参数，防止目录遍历攻击
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	
	// 检查路径是否存在
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", errors.New("path does not exist")
	}
	
	// 检查git命令是否存在
	if _, err := exec.LookPath("git"); err != nil {
		return "", errors.New("git command not found")
	}
	
	// 使用带超时控制的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cmd := exec.CommandContext(ctx, "git", "diff")
	cmd.Dir = absPath

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// parseGitDiff 解析git diff输出，提取更改的文件
func parseGitDiff(diff string) []string {
	var changedFiles []string
	lines := strings.Split(diff, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git") {
			// 提取文件路径
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				// 提取新文件路径（去掉a/和b/前缀）
				filePath := strings.TrimPrefix(parts[3], "b/")
				if strings.HasSuffix(filePath, ".go") {
					changedFiles = append(changedFiles, filePath)
				}
			}
		}
	}

	return changedFiles
}

// getChangedGoFiles 获取更改的Go文件
func getChangedGoFiles(path string) ([]string, error) {
	diff, err := getGitDiff(path)
	if err != nil {
		return nil, err
	}

	return parseGitDiff(diff), nil
}
