package main

import (
	"log/slog"
	"os"
	"log"
	"context"
)


func main() {
	var t Test
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Test logger")
	logger.InfoContext(context.Background(), "Test logger")

	
	
	t.test()

	logs := log.New(os.Stdout, "", 0)
	logs.Print("Test")
}


type Test struct {}
func (t Test) test() {}

