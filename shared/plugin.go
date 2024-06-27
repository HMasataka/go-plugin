package shared

import (
	"net/rpc"

	plugin "github.com/hashicorp/go-plugin"
)

type Greeter interface {
	Greet() string
}

type GreeterRPC struct {
	Client *rpc.Client
}

func (g *GreeterRPC) Greet() string {
	var resp string

	err := g.Client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(_ interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

func NewPlugin(impl Greeter) plugin.Plugin {
	return &GreeterPlugin{
		Impl: impl,
	}
}

type GreeterPlugin struct {
	Impl Greeter
}

func (m GreeterPlugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: m.Impl}, nil
}

func (GreeterPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{Client: c}, nil
}
