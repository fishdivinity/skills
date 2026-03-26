package analyzer

import (
	"os"
	"path/filepath"
)

// findGoFiles 查找指定路径下的所有Go文件
func findGoFiles(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(filePath) == ".go" {
			files = append(files, filePath)
		}

		return nil
	})

	return files, err
}
