package internal
import (
	"go/ast"
)

type targetLog string

const slogLog = targetLog("*log/slog.Logger")
const zapLog = targetLog("*go.uber.org/zap.Logger")

type LogCall struct {
    Package targetLog 
    Method    string 
    Call      *ast.CallExpr
    MsgIndex  int
}