package App

import (
	"cloudbees_assignmnet.com/ass1/internal/App/local_repo"
	"cloudbees_assignmnet.com/ass1/internal/App/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"App",
	local_repo.Module,
	service.Module,
)