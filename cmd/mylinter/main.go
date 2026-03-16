package main

import (
	"aartchik.mylinter/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
	"flag"
	"aartchik.mylinter/internal"
)


func main() {
	configPath := flag.String("config", ".mylinter_sensitive.yml", "path to sensitive words file")
	flag.Parse()

	internal.LoadConfig(*configPath)
	singlechecker.Main(analyzer.Analyzer)
}