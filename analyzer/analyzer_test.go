package analyzer

import (
	"testing"
	"golang.org/x/tools/go/analysis/analysistest"
)


func TestAnalyzerCleanSlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "clean_slog")
}

func TestAnalyzerCleanZap(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "clean_zap")
}

func TestAnalyzerDirtySlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "dirty_slog")
}

func TestAnalyzerDirtyZap(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "dirty_zap")
}
func TestFix(t *testing.T) {
	analysistest.RunWithSuggestedFixes( t, analysistest.TestData(), Analyzer, "runWithSuggestedFixes")
}