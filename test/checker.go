package main

import (
	"go.uber.org/zap"

)


func main() {
	zap.L().Info("message", zap.Bool("string", true))
}