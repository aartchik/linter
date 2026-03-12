package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

type targetLog string

const slogLog = targetLog("*log/slog.Logger")
const zapLog = targetLog("*go.uber.org/zap.Logger")

func targetLogCall(pass *analysis.Pass, call *ast.CallExpr) (targetLog, string) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", ""
	}

	tv, ok := pass.TypesInfo.Types[sel.X]
	if !ok || tv.Type == nil {
		return "", ""
	}

	typeStr := targetLog(tv.Type.String())
	switch typeStr {
	case slogLog:
		return slogLog, sel.Sel.Name
	case zapLog:
		return zapLog, sel.Sel.Name
	}

	return "", ""
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

func isEnglish(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			return false
		}
	}
	return true
}

func hasSpecialSymbols(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			continue
		}
		return true
	}
	return false
}



func checkArgs(idx int, arg string) bool {
	arg = strings.TrimSpace(arg)
	r := []rune(arg)

	if len(r) == 0 {
		return true
	}


	if idx == 0 || idx%2 == 1 {
		if unicode.IsLetter(r[0]) && unicode.IsUpper(r[0]) {
			return false
		}
	}
	return true
}


func linterSLOG(pass *analysis.Pass, call *ast.CallExpr) (bool) {

	target, method := targetLogCall(pass, call)
	if target == "" {
		return true
	}
	startMsg := 0
	if strings.HasSuffix(method, "Context") {
		startMsg = 1
	}

	if len(call.Args) <= startMsg {
		return true
	}

	msgParts := collectStrings(call.Args[startMsg])
	for _, part := range msgParts {
		if !checkArgs(0, part) {
			pass.Reportf(call.Args[startMsg].Pos(), "log message should start with lowercase")
			return true
		}
		if !isEnglish(part) {
			pass.Reportf(call.Args[startMsg].Pos(), "log message should contains only english letters")
			return true
		}
		if hasSpecialSymbols(part) {
			pass.Reportf(call.Args[startMsg].Pos(), "log message contains special symbols or emoji")
			return true
		} 
	}

	fieldIdx := 1
	for i := startMsg + 1; i < len(call.Args); i++ {
		arg := call.Args[i]

		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			if (i-(startMsg+1))%2 == 0 {
				for _, s := range collectStrings(arg) {
					if !checkArgs(fieldIdx, s) {
						pass.Reportf(arg.Pos(), "log field key should start with lowercase")
						return true
					}
					if !isEnglish(s) {
						pass.Reportf(call.Args[startMsg].Pos(), "log message should contains only english letters")
						return true
					}
					if hasSpecialSymbols(s) {
						pass.Reportf(call.Args[startMsg].Pos(), "log message contains special symbols or emoji")
						return true
					} 
				}
				fieldIdx += 2
			}
			continue
		}

		if innerCall, ok := arg.(*ast.CallExpr); ok {
			parts := collectStrings(innerCall)
			if len(parts) > 0 {
				key := parts[0]
				if !checkArgs(fieldIdx, key) {
					pass.Reportf(arg.Pos(), "log field key should start with lowercase")
					return true
				}
				if !isEnglish(key) {
					pass.Reportf(call.Args[startMsg].Pos(), "log filed key should contains only english letters")
					return true
				}
				if hasSpecialSymbols(key) {
					pass.Reportf(call.Args[startMsg].Pos(), "log filed key contains special symbols or emoji")
					return true
				} 
				fieldIdx += 2
			}
		}
	}

	return true
}


func linterZAP(pass *analysis.Pass, call *ast.CallExpr) (bool) { return true }

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			target, _ := targetLogCall(pass, call)
			switch  target{

			case "":
				return true
			case slogLog:
				return linterSLOG(pass, call)
			case zapLog:
				return linterZAP(pass, call)
			}

			return true
		})
	}

	return nil, nil
}