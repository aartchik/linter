package internal

import (
	"os"

	"gopkg.in/yaml.v3"
	"strings"
)

var DefaultSensitiveWordsInKey = map[string]struct{}{
	"password": {},
	"passwd": {},
	"pwd": {},
	"token": {},
	"secret": {},
	"apikey": {},
	"api_key": {},
	"access_token": {},
	"refresh_token": {},
	"authorization": {},
}

type Config struct {
	SensitivePatt []string `yaml:"sensitive_patterns"`
}


var SensitivePatterns = map[string]struct{}{}
var config Config

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return
	}


	for _, p := range cfg.SensitivePatt {
		SensitivePatterns[strings.ToLower(p)] = struct{}{}
	}
}


