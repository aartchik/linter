package main

import (
	"log/slog"
	"os"
	"context"
	//"go.uber.org/zap"
)


func main() {
	var t Test
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("1Test logger", "field1", "22test1")
	logger.InfoContext(context.WithValue(context.Background(), "Test2", "val2"), "test2 logger")
	logger.Info("test3 logger!!!", "userID", "5")



	t.test()
	//logs, _ := zap.NewProduction()
}


type Test struct {}
func (t Test) test() {}

