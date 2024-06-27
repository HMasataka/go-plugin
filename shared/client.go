package shared

import (
	"errors"
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type PluginRPC struct {
	Client *rpc.Client
}

func (g *PluginRPC) Greet(args string) (string, error) {
	var resp string

	err := g.Client.Call("Plugin.Greet", args, &resp)
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
