package analyzer

import "golang.org/x/tools/go/analysis"

var	Analyzer = &analysis.Analyzer{
	Name: "logrecords",
	Doc:  "check for correct format log records",
	Run:  run,
}