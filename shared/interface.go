package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Greeter interface {
	Greet() string
}

type GreeterPlugin struct {
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
