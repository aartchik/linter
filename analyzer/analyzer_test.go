package analyzer

import (
	"testing"
	"golang.org/x/tools/go/analysis/analysistest"
	"aartchik.mylinter/internal"
)


func TestAnalyzerCleanSlog(t *testing.T) {
	internal.SensitivePatterns = internal.DefaultSensitiveWordsInKey
	analysistest.Run(t, analysistest.TestData(), Analyzer, "clean_slog")
}

func TestAnalyzerCleanZap(t *testing.T) {
	internal.SensitivePatterns = internal.DefaultSensitiveWordsInKey
	analysistest.Run(t, analysistest.TestData(), Analyzer, "clean_zap")
}

func TestAnalyzerDirtySlog(t *testing.T) {
	internal.SensitivePatterns = internal.DefaultSensitiveWordsInKey
	analysistest.Run(t, analysistest.TestData(), Analyzer, "dirty_slog")
}

func TestAnalyzerDirtyZap(t *testing.T) {
	internal.SensitivePatterns = internal.DefaultSensitiveWordsInKey
	analysistest.Run(t, analysistest.TestData(), Analyzer, "dirty_zap")
}
func TestFix(t *testing.T) {
	internal.SensitivePatterns = internal.DefaultSensitiveWordsInKey
	analysistest.RunWithSuggestedFixes( t, analysistest.TestData(), Analyzer, "run_with_suggested_fixes")
}