package config

import (
	"encoding/json"
	"golang.org/x/exp/slog"
	"strconv"
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
	parseBool, err := strconv.ParseBool(env)
	if err != nil {
		parseBool = false
		slog.Error("Impossible to parse", "error", err.Error())
	}
	d := &DefaultConfig{TracingEnabled: parseBool, TypeOfConfig: "default"}
	slog.Info("Running eventstore-grpc config:", "shape-app-config", d.Print())
	return d
}
