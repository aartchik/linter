package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"aartchik.mylinter/analyzer"
)

func init() {
	register.Plugin("mylinter", New)
}

type Settings struct{}

type Plugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	_, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}