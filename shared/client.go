package shared

import "net/rpc"

type GreeterRPC struct {
	client *rpc.Client
}

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}
