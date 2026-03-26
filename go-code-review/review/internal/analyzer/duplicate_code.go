package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// detectDuplicateCode 检测重复代码模式
func detectDuplicateCode(filePath string) ([]Issue, error) {
	var issues []Issue

	// 解析Go文件
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return issues, err
	}

	// 检测切片遍历检查元素是否存在的模式
	detectContainsPattern(node, fset, filePath, &issues)

	return issues, nil
}

// detectContainsPattern 检测切片遍历检查元素是否存在的模式
func detectContainsPattern(node ast.Node, fset *token.FileSet, filePath string, issues *[]Issue) {
	ast.Inspect(node, func(n ast.Node) bool {
		// 查找if语句
		ifStmt, ok := n.(*ast.IfStmt)
		if !ok {
			return true
		}

		// 检查if语句体
		if len(ifStmt.Body.List) == 0 {
			return true
		}

		// 查找for循环
		forStmt, ok := ifStmt.Body.List[0].(*ast.ForStmt)
		if !ok {
			return true
		}

		// 检查是否是range循环
		if forStmt.Init != nil || forStmt.Post != nil {
			return true
		}

		// 检查循环体是否设置变量为true
		hasSetTrue := false
		var varName string

		ast.Inspect(forStmt.Body, func(n ast.Node) bool {
			// 查找赋值语句
			assign, ok := n.(*ast.AssignStmt)
			if !ok || len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
				return true
			}

			// 检查左边是否是标识符
			lhsIdent, ok := assign.Lhs[0].(*ast.Ident)
			if !ok {
				return true
			}

			// 检查右边是否是true
			// 方式1：作为标识符（最常见的情况）
			ident, ok := assign.Rhs[0].(*ast.Ident)
			if ok && ident.Name == "true" {
				varName = lhsIdent.Name
				hasSetTrue = true
				return false
			}

			// 方式2：作为基本字面量（较少见）
			rhsBool, ok := assign.Rhs[0].(*ast.BasicLit)
			if !ok || rhsBool.Kind != token.STRING || rhsBool.Value != "true" {
				return true
			}

			varName = lhsIdent.Name
			hasSetTrue = true
			return false
		})

		if hasSetTrue && varName != "" {
			pos := fset.Position(ifStmt.Pos())
			*issues = append(*issues, Issue{
				Type:        "DuplicateCode",
				Description: "Possible duplicate pattern: slice contains check",
				FilePath:    filePath,
				Line:        pos.Line,
				Severity:    "Low",
			})
		}

		return true
	})
}
