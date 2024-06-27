package shared

import (
	"errors"
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type PluginRPC struct {
	Client *rpc.Client
}

func (g *PluginRPC) Greet() string {
	var resp string

	err := g.Client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

type GreeterPluginClient struct {
	Impl Plugin
}

func (GreeterPluginClient) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("Not implemented")
}

func (GreeterPluginClient) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPC{Client: c}, nil
}
