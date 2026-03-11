package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"golang.org/x/tools/go/analysis"
	"unicode"
)

type targetLog string
const slogLog = targetLog("*log/slog.Logger")
const logLog = targetLog("*log.Logger")

const zapLog = targetLog("log/slog.Logger")



func TargetLogCall(pass *analysis.Pass, call *ast.CallExpr) targetLog {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	_, ok = sel.X.(*ast.Ident)
	if !ok {
		return ""
	}

	tv, ok := pass.TypesInfo.Types[sel.X]
	if !ok || tv.Type == nil {
		return ""
	}

	typeStr := targetLog(tv.Type.String())
	switch typeStr {

	case slogLog:
		return slogLog
	
	case logLog:
		return logLog
	
	case zapLog:
		return zapLog
	}
	return ""
	
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

func checkArgs(args []string) bool {
	for _, item := range args {
		r := []rune(item)
		if len(r) > 0 && unicode.IsUpper(r[0]) {
			return false
		}
	}
	return true
}


func run(pass *analysis.Pass) (any, error) {

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			} 
			if target := TargetLogCall(pass, call); target == "" {
				return true
			}	 
			args := collectStrings(call.Args[0])
			ok = checkArgs(args)
			if !ok {
				pass.Reportf(n.Pos(), fmt.Sprintf("func called: %s", call.Fun.(*ast.SelectorExpr).Sel.Name))
			}
			return true
		})
	}
	return nil, nil
}