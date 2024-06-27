package shared

import (
	plugin "github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "plugin",
	MagicCookieValue: "master_data_converter",
}
