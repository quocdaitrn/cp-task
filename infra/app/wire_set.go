package app

import (
	"github.com/google/wire"
	"github.com/quocdaitrn/golang-kit/validator"

	"github.com/quocdaitrn/cp-task/domain/service/serviceimpl"
	"github.com/quocdaitrn/cp-task/infra/adapters"
	"github.com/quocdaitrn/cp-task/infra/config"
	"github.com/quocdaitrn/cp-task/infra/providers"
	"github.com/quocdaitrn/cp-task/infra/repo/rpcimpl"
	"github.com/quocdaitrn/cp-task/infra/repo/storeimpl"
)

var ApplicationSet = wire.NewSet(
	config.ProvideConfig,
	validator.New,

	adapters.ProvideMySQL,
	adapters.ProvideRoutes,
	adapters.ProvideRestService,
	providers.ProvideLogger,
	adapters.ProvideGRPCAuthClient,
	adapters.ProvideGRPCUserServiceClient,

	storeimpl.NewTaskRepo,
	rpcimpl.NewUserRepo,
	serviceimpl.NewTaskService,
)
