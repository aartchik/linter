package internal

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)


func targetLogCall(pass *analysis.Pass, call *ast.CallExpr) LogCall {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return LogCall{}
	}
	

	if ident, ok := sel.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkgName, ok := obj.(*types.PkgName); ok && pkgName.Imported().Path() == "log/slog" {
				if !isSupportedMethodSlog(sel.Sel.Name) {
					return LogCall{}
				}
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
		if ident, ok := sel.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkgName, ok := obj.(*types.PkgName); ok && pkgName.Imported().Path() == "go.uber.org/zap" {
				if !isSupportedMethodZap(sel.Sel.Name) {
					return LogCall{}
				}
				log := LogCall{
					Package: zapLog,
					Method:  sel.Sel.Name,
					Call:    call,
					MsgIndex: 0,
				}
				if log.Method == "Log" {
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
		if !isSupportedMethodSlog(log.Method) {
			return LogCall{}
		}
		if strings.HasSuffix(sel.Sel.Name, "Context") {
			log.MsgIndex = 1
		}
		return log
	case zapLog:
		log := LogCall{Package: zapLog, Method: sel.Sel.Name, Call: call, MsgIndex: 0}
		if !isSupportedMethodZap(log.Method) {
			return LogCall{}
		}
		if log.Method == "Log" {
			log.MsgIndex = 1
		}
		return log
	}

	return LogCall{Package: "", Method: "", Call: nil, MsgIndex: 0}
}