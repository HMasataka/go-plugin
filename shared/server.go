package shared

import (
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

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
