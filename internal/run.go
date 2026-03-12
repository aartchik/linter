package internal
import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func Run(pass *analysis.Pass) (any, error) {
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