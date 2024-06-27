package main

import (
	"os"

	"github.com/HMasataka/go-plugin/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	greeter := &GreeterFoo{
		logger: logger,
	}

	var pluginMap = map[string]plugin.Plugin{
		"foo": &shared.GreeterPlugin{Impl: greeter},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
