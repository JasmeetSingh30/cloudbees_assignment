package server

import "go.uber.org/fx"

var Module = fx.Module(
	"Server Service Mapper",
	fx.Provide(
		NewServerServiceMapper,
	),
)
