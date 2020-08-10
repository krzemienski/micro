// Package auth is an authentication plugin
package auth

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/server"
	"github.com/micro/micro/v3/internal/wrapper"
	"github.com/micro/micro/v3/plugin"
	c "github.com/micro/micro/v3/service/client"
	s "github.com/micro/micro/v3/service/server"
)

var (
	Plugin = plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithInit(func(ctx *cli.Context) error {
			// wrap the client
			c.DefaultClient = wrapper.AuthClient(c.DefaultClient)

			// wrap the server
			s.DefaultServer.Init(
				server.WrapHandler(wrapper.AuthHandler()),
			)
			return nil
		}),
	)
)

func init() {
	plugin.Register(Plugin)
}