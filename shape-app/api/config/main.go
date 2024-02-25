package config

import (
	"encoding/json"
	"golang.org/x/exp/slog"
)

type Tracing interface {
	IsTracingEnabled() bool
}

type Observability interface {
	Tracing
}
type Config interface {
	Observability
	Print() string
}

type DefaultConfig struct {
	TracingEnabled bool   `json:"tracingEnabled"`
	TypeOfConfig   string `json:typeOfConfig"`
}

func (d *DefaultConfig) Print() string {
	marshal, _ := json.Marshal(d)
	return string(marshal)
}

func (d *DefaultConfig) IsTracingEnabled() bool {
	return d.TracingEnabled
}

func GetConfigFrom(env string) Config {
	d := &DefaultConfig{TracingEnabled: false, TypeOfConfig: "default"}
	slog.Info("Running shape-app config:", "shape-app-config", d.Print())
	return d
}
