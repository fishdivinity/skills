package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// detectMapPreallocation 检测map预分配
func detectMapPreallocation(filePath string) ([]Issue, error) {
	var issues []Issue

	// 解析Go文件
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return issues, err
	}

	// 遍历AST查找map创建
	ast.Inspect(node, func(n ast.Node) bool {
		// 查找make函数调用
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		// 检查是否是make函数
		ident, ok := call.Fun.(*ast.Ident)
		if !ok || ident.Name != "make" {
			return true
		}

		// 检查参数数量
		if len(call.Args) < 1 {
			return true
		}

		// 检查是否是map类型
		_, ok = call.Args[0].(*ast.MapType)
		if !ok {
			return true
		}

		// 检查是否预分配了容量
		if len(call.Args) == 1 {
			// 没有预分配容量
			pos := fset.Position(call.Pos())
			issues = append(issues, Issue{
				Type:        "MapPreallocation",
				Description: "Map created without preallocation",
				FilePath:    filePath,
				Line:        pos.Line,
				Severity:    "Low",
			})
		}

		return true
	})

	return issues, nil
}
