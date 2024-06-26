//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/HMasataka/go-plugin/greeting"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	greeting.RegisterGreeter(MyPlugin{})
}

type MyPlugin struct{}

func (m MyPlugin) SayHello(ctx context.Context, request greeting.GreetRequest) (greeting.GreetReply, error) {
	return greeting.GreetReply{
		Message: "Hello, " + request.GetName(),
	}, nil
}
