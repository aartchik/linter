package internal

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

func linterSLOG(pass *analysis.Pass, log *LogCall) {
	if len(log.Call.Args) <= log.MsgIndex {
		return
	}

	msgArg := log.Call.Args[log.MsgIndex]
	checkMessage(pass, msgArg)

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		arg := log.Call.Args[i]
		offset := i - (log.MsgIndex + 1)

		if offset%2 == 0 {
			checkSlogKey(pass, arg)
		}
	}

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		checkSensitiveArg(pass, log.Call.Args[i])
	}
}

func checkMessage(pass *analysis.Pass, msgArg ast.Expr) {
	msgParts := collectStrings(msgArg)
	if len(msgParts) == 0 {
		return
	}

	msg := msgParts[0]

	if !checkLowerCase(msg) {
		newMsg := toLowerCase(msg)
		pass.Report(analysis.Diagnostic{
			Pos:     msgArg.Pos(),
			End:     msgArg.End(),
			Message: "log message should start with lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     msgArg.Pos(),
							End:     msgArg.End(),
							NewText: []byte(strconv.Quote(newMsg)),
						},
					},
				},
			},
		})
	}

	if !isEnglish(msg) {
		pass.Reportf(msgArg.Pos(), "log message should contain only English letters")
	}

	if !notHasSpecialSymbols(msg) {
		newMsg := toStandardSymbols(msg)
		pass.Report(analysis.Diagnostic{
			Pos:     msgArg.Pos(),
			End:     msgArg.End(),
			Message: "log message contains special symbols or emoji",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     msgArg.Pos(),
							End:     msgArg.End(),
							NewText: []byte(strconv.Quote(newMsg)),
						},
					},
				},
			},
		})
	}

	if !notContainsSensitiveWordInMsg(msg) {
		pass.Reportf(msgArg.Pos(), "log message contains potential sensitive word")
	}
}

func checkSlogKey(pass *analysis.Pass, arg ast.Expr) {
	switch v := arg.(type) {
	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return
		}
		checkKeyString(pass, v, v)

	case *ast.CallExpr:
		if len(v.Args) == 0 {
			return
		}

		keyArg := v.Args[0]

		lit, ok := keyArg.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}

		checkKeyString(pass, keyArg, lit)
	}
}

func checkKeyString(pass *analysis.Pass, reportNode ast.Expr, strExpr ast.Expr) {
	strs := collectStrings(strExpr)
	if len(strs) == 0 {
		return
	}

	key := strs[0]

	if !checkLowerCase(key) {
		newKey := toLowerCase(key)
		pass.Report(analysis.Diagnostic{
			Pos:     reportNode.Pos(),
			End:     reportNode.End(),
			Message: "log field key should start with lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     strExpr.Pos(),
							End:     strExpr.End(),
							NewText: []byte(strconv.Quote(newKey)),
						},
					},
				},
			},
		})
	}

	if !isEnglish(key) {
		pass.Reportf(reportNode.Pos(), "log field key should contain only English letters")
	}

	if !notHasSpecialSymbols(key) {
		newKey := toStandardSymbols(key)
		pass.Report(analysis.Diagnostic{
			Pos:     reportNode.Pos(),
			End:     reportNode.End(),
			Message: "log field key contains special symbols or emoji",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     strExpr.Pos(),
							End:     strExpr.End(),
							NewText: []byte(strconv.Quote(newKey)),
						},
					},
				},
			},
		})
	}
}

func checkSensitiveArg(pass *analysis.Pass, arg ast.Expr) {
	ast.Inspect(arg, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if isSensitiveKey(x.Name) {
				pass.Reportf(x.Pos(), "log field key contains potential sensitive word")
				return false
			}

		case *ast.BasicLit:
			if x.Kind != token.STRING {
				return true
			}
			for _, s := range collectStrings(x) {
				if isSensitiveKey(s) {
					pass.Reportf(x.Pos(), "log field key contains potential sensitive word")
					return false
				}
			}
		}
		return true
	})
}