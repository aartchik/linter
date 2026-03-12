package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"
	"go/types"
	"golang.org/x/tools/go/analysis"
)


type LogCall struct {
    Package targetLog 
    Method    string 
    Call      *ast.CallExpr
    MsgIndex  int
}

type targetLog string

const slogLog = targetLog("*log/slog.Logger")
const zapLog = targetLog("*go.uber.org/zap.Logger")

func targetLogCall(pass *analysis.Pass, call *ast.CallExpr) LogCall {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return LogCall{}
	}
	if !isSupportedMethodSlog(sel.Sel.Name) {
		return LogCall{}
	}

	if ident, ok := sel.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkgName, ok := obj.(*types.PkgName); ok && pkgName.Imported().Path() == "log/slog" {
				log := LogCall{
					Package: slogLog,
					Method:  sel.Sel.Name,
					Call:    call,
					MsgIndex: 0,
				}
				if strings.HasSuffix(sel.Sel.Name, "Context") {
					log.MsgIndex = 1
				}
				return log
			}
		}
	}


	tv, ok := pass.TypesInfo.Types[sel.X]
	if !ok || tv.Type == nil {
		return LogCall{}
	}

	typeStr := targetLog(tv.Type.String())
	switch typeStr {
	case slogLog:
		log := LogCall{Package: slogLog, Method: sel.Sel.Name, Call: call}
		if strings.HasSuffix(sel.Sel.Name, "Context") {
			log.MsgIndex = 1
		}
		return log
	case zapLog:
		return LogCall{Package: zapLog, Method: sel.Sel.Name, Call: call, MsgIndex: 0}
	}

	return LogCall{Package: "", Method: "", Call: nil, MsgIndex: 0}
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
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '_' || r == '-'{
			continue
		}
		return true
	}
	return false
}

func isSupportedMethodSlog(method string) bool {
	switch method {
	case "Debug", "Info", "Warn", "Error", "DebugContext", "InfoContext", "WarnContext", "ErrorContext":
		return true
	default:
		return false
	}
}

func checkLowercase(arg string) bool {
	arg = strings.TrimSpace(arg)
	r := []rune(arg)

	if len(r) == 0 {
		return true
	}

	return !(unicode.IsLetter(r[0]) && unicode.IsUpper(r[0]))
}


func linterSLOG(pass *analysis.Pass, log *LogCall) {
	if len(log.Call.Args) <= log.MsgIndex {
		return
	}

	msgArg := log.Call.Args[log.MsgIndex]
	msgParts := collectStrings(msgArg)

	for _, part := range msgParts {
		if !checkLowercase(part) {
			pass.Reportf(msgArg.Pos(), "log message should start with lowercase")
		}
		if !isEnglish(part) {
			pass.Reportf(msgArg.Pos(), "log message should contain only English letters")
		}
		if hasSpecialSymbols(part) {
			pass.Reportf(msgArg.Pos(), "log message contains special symbols or emoji")
		}
	}

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		arg := log.Call.Args[i]

		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			if (i-(log.MsgIndex+1))%2 == 0 {
				for _, s := range collectStrings(arg) {
					if !checkLowercase(s) {
						pass.Reportf(arg.Pos(), "log field key should start with lowercase")
					}
					if !isEnglish(s) {
						pass.Reportf(arg.Pos(), "log field key should contain only English letters")
					}
					if hasSpecialSymbols(s) {
						pass.Reportf(arg.Pos(), "log field key contains special symbols or emoji")
					}
				}
			}
			continue
		}

		if innerCall, ok := arg.(*ast.CallExpr); ok {
			if len(innerCall.Args) == 0 {
				continue
			}

			keyParts := collectStrings(innerCall.Args[0])
			if len(keyParts) == 0 {
				continue
			}

			key := keyParts[0]
			if !checkLowercase(key) {
				pass.Reportf(arg.Pos(), "log field key should start with lowercase")
			}
			if !isEnglish(key) {
				pass.Reportf(arg.Pos(), "log field key should contain only English letters")
			}
			if hasSpecialSymbols(key) {
				pass.Reportf(arg.Pos(), "log field key contains special symbols or emoji")
			}
		}
	}
}


func linterZAP(pass *analysis.Pass, log *LogCall) { }

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			log := targetLogCall(pass, call)
			switch  log.Package {

			case "":
				return true
			case slogLog:
				linterSLOG(pass, &log)
				return true
			case zapLog:
				linterZAP(pass, &log)
				return true
			}

			return true
		})
	}

	return nil, nil
}