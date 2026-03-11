package main

import (
	"log/slog"
	"os"
	"log"
)


func main() {
	var t Test
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("test logger")
	
	t.test()

	logs := log.New(os.Stdout, "", 0)
	logs.Print("test")
}


type Test struct {}
func (t Test) test() {}

