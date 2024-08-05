package local_repo

import "go.uber.org/fx"

var Module = fx.Module(
	"Local Repository",
	fx.Provide(
		NewTicketRepository,
		NewPassengerRepository,
		NewTrainRepository,
	),
)
