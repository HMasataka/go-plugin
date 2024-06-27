package shared

import (
	"errors"
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type Plugin interface {
	Greet() string
}

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

type PluginRPCServer struct {
	Impl Plugin
}

func (s *PluginRPCServer) Greet(_ interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

func NewPluginServer(impl Plugin) plugin.Plugin {
	return &GreeterPluginServer{
		Impl: impl,
	}
}

type GreeterPluginServer struct {
	Impl Plugin
}

func (m GreeterPluginServer) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &PluginRPCServer{Impl: m.Impl}, nil
}

func (GreeterPluginServer) Client(_ *plugin.MuxBroker, _ *rpc.Client) (interface{}, error) {
	return nil, errors.New("Not implemented")
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
