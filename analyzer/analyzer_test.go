package analyzer

import (
	"testing"
	"golang.org/x/tools/go/analysis/analysistest"
)


func TestAnalyzerCleanSlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "clean_slog")
}

func TestAnalyzerDirtySlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "dirty_slog")
}