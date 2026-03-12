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
	msgParts := collectStrings(msgArg)
	if !checkLowerCase(msgParts[0]) {
		newmsg := toLowerCase(msgParts[0])
		pass.Report(analysis.Diagnostic {
			Pos: msgArg.Pos(),
			End: msgArg.End(),
			Message: "log message should start with lowercase",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					TextEdits: []analysis.TextEdit{
						{
							Pos:     msgArg.Pos(),
							End:     msgArg.End(),
							NewText: []byte(strconv.Quote(newmsg)),
						},
					},
				},
			},
	})
}
	

	for _, part := range msgParts {

		if !isEnglish(part) {
			pass.Reportf(msgArg.Pos(), "log message should contain only English letters")
		}
		if !notHasSpecialSymbols(part) {
			pass.Reportf(msgArg.Pos(), "log message contains special symbols or emoji")
		}
		if !notContainsSensitiveWord(part) {
			pass.Reportf(msgArg.Pos(), "log field key contains potential sensitive word")
		}
	}

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		arg := log.Call.Args[i]

		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			if (i-(log.MsgIndex+1))%2 == 0 {
				for _, s := range collectStrings(arg) {
					if !checkLowerCase(s) {
						pass.Reportf(arg.Pos(), "log field key should start with lowercase")
					}
					if !isEnglish(s) {
						pass.Reportf(arg.Pos(), "log field key should contain only English letters")
					}
					if !notHasSpecialSymbols(s) {
						pass.Reportf(arg.Pos(), "log field key contains special symbols or emoji")
					}
					if !notContainsSensitiveWord(s) {
						pass.Reportf(arg.Pos(), "log field key contains potential sensitive word")
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
			if !checkLowerCase(key) {
				pass.Reportf(arg.Pos(), "log field key should start with lowercase")
			}
			if !isEnglish(key) {
				pass.Reportf(arg.Pos(), "log field key should contain only English letters")
			}
			if !notHasSpecialSymbols(key) {
				pass.Reportf(arg.Pos(), "log field key contains special symbols or emoji")
			}
			if !notContainsSensitiveWord(key) {
				pass.Reportf(arg.Pos(), "log field key contains potential sensitive word")
			}
		}
	}
}