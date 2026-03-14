package internal

import (
	"go/ast"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

func linterZAP(pass *analysis.Pass, log *LogCall) {
	if len(log.Call.Args) <= log.MsgIndex {
		return
	}

	msgArg := log.Call.Args[log.MsgIndex]
	msgParts := collectStrings(msgArg)

	if len(msgParts) > 0 && !checkLowerCase(msgParts[0]) {
		newmsg := toLowerCase(msgParts[0])
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
			newmsg := toStandardSymbols(part)
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
								NewText: []byte(strconv.Quote(newmsg)),
							},
						},
					},
				},
			})
		}
		if !notContainsSensitiveWordInMsg(part) {
			pass.Reportf(msgArg.Pos(), "log message contains potential sensitive word")
		}
	}

	for i := log.MsgIndex + 1; i < len(log.Call.Args); i++ {
		arg := log.Call.Args[i]

		innerCall, ok := arg.(*ast.CallExpr)
		if !ok {
			continue
		}


		if len(innerCall.Args) == 0 {
			continue
		}

		keyArg := innerCall.Args[0]
		keyParts := collectStrings(keyArg)
		if len(keyParts) == 0 {
			continue
		}

		key := keyParts[0]

		if !checkLowerCase(key) {
			newmsg := toLowerCase(key)
			pass.Report(analysis.Diagnostic{
				Pos:     keyArg.Pos(),
				End:     keyArg.End(),
				Message: "log field key should start with lowercase",
				SuggestedFixes: []analysis.SuggestedFix{
					{
						TextEdits: []analysis.TextEdit{
							{
								Pos:     keyArg.Pos(),
								End:     keyArg.End(),
								NewText: []byte(strconv.Quote(newmsg)),
							},
						},
					},
				},
			})
		}

		if !isEnglish(key) {
			pass.Reportf(keyArg.Pos(), "log field key should contain only English letters")
		}

		if !notHasSpecialSymbols(key) {
			newmsg := toStandardSymbols(key)
			pass.Report(analysis.Diagnostic{
				Pos:     keyArg.Pos(),
				End:     keyArg.End(),
				Message: "log field key contains special symbols or emoji",
				SuggestedFixes: []analysis.SuggestedFix{
					{
						TextEdits: []analysis.TextEdit{
							{
								Pos:     keyArg.Pos(),
								End:     keyArg.End(),
								NewText: []byte(strconv.Quote(newmsg)),
							},
						},
					},
				},
			})
		}

		if isSensitiveKey(key) {
			pass.Reportf(keyArg.Pos(), "log field key contains potential sensitive word")
		}

		for _, valueArg := range innerCall.Args[1:] {
			idents := collectIdents(valueArg)
			for _, ident := range idents {
				obj := pass.TypesInfo.Uses[ident]
				if obj == nil {
					continue
				}

				if _, ok := obj.(*types.Var); !ok {
					continue
				}

				if isSensitiveKey(ident.Name) {
					pass.Reportf(ident.Pos(), "log field key contains potential sensitive word")
				}
			}
		}
	}
}