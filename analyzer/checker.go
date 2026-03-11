package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"golang.org/x/tools/go/analysis"
)

type targetLog string
const slogLog = targetLog("*log/slog.Logger")
const logLog = targetLog("*log.Logger")

const zapLog = targetLog("log/slog.Logger")



func isTargetLogCall(pass *analysis.Pass, call *ast.CallExpr) (bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	_, ok = sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	tv, ok := pass.TypesInfo.Types[sel.X]
	if !ok || tv.Type == nil {
		return false
	}

	typeStr := tv.Type.String()
	if typeStr == string(slogLog) || typeStr == string(logLog) {
		fmt.Printf("type: %s", typeStr)
		return true
	}
	//pass.Reportf(n.Pos(), fmt.Sprintf("func called: %s", sel.Sel.Name))
	return true
}

func collectStrings(expr ast.Expr) []string {
	var result []string

	switch v := expr.(type) {

	case *ast.BasicLit:
		if v.Kind == token.STRING {
			s, err := strconv.Unquote(v.Value)
			if err == nil {
				result = append(result, s)
			}
		}

	case *ast.BinaryExpr:
		if v.Op == token.ADD {
			result = append(result, collectStrings(v.X)...)
			result = append(result, collectStrings(v.Y)...)
		}

	case *ast.CallExpr:
		for _, arg := range v.Args {
			result = append(result, collectStrings(arg)...)
		}

	case *ast.ParenExpr:
		result = append(result, collectStrings(v.X)...)

	}
	return result
}


func run(pass *analysis.Pass) (any, error) {

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			} 
			ok = isTargetLogCall(pass, call)
			if !ok {
				return false
			}
			return true
			
		})
	}
	return nil, nil
}