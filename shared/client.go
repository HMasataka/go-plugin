package shared

import (
	"errors"
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type PluginRPC struct {
	Client *rpc.Client
}

func (g *PluginRPC) Greet(args1, args2 string) (string, error) {
	var resp string

	params := &Arg{P1: args1, P2: args2}

	err := g.Client.Call("Plugin.Greet", params, &resp)
	if err != nil {
		return "", err
	}

	return resp, nil
}

func NewGreeterPluginClient() plugin.Plugin {
	return &GreeterPluginClient{}
}

type GreeterPluginClient struct{}

func (GreeterPluginClient) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("Not implemented")
}

func (GreeterPluginClient) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPC{Client: c}, nil
}
