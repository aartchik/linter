package main

import (
	"context"
	"log/slog"
	"os"

	"go.uber.org/zap"
	//"go.uber.org/zap"
)


func main() {
	var t Test
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("1Test logger", "Field1", "22test1")
	logger.InfoContext(context.WithValue(context.Background(), "Test2", "val2"), "test2 logger")
	logger.Info("test3 logger!!!", "userID", "5")
	slog.Info("Test")
	zap.L().Info("test")


	t.test()
	//logs, _ := zap.NewProduction()
}


type Test struct {}
func (t Test) test() {}

