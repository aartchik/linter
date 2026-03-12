package internal

import (
	"go/ast"
	"go/token"
	"strconv"
)

func collectIdents(expr ast.Expr) []*ast.Ident {
	var result []*ast.Ident

	ast.Inspect(expr, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		result = append(result, ident)
		return true
	})
	return result
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