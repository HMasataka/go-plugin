package main

import (
	"github.com/hashicorp/go-hclog"
)

type GreeterFoo struct {
	logger hclog.Logger
}

func (g *GreeterFoo) Greet() string {
	g.logger.Debug("message from GreeterHello.Greet")
	return "Hello!"
}
