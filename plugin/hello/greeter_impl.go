package main

import (
	"github.com/hashicorp/go-hclog"
)

type GreeterHello struct {
	logger hclog.Logger
}

func (g *GreeterHello) Greet(msg1, msg2 string) (string, error) {
	g.logger.Debug("message from GreeterHello.Greet")
	return msg1 + msg2, nil
}
