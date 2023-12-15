//go:build wireinject
// +build wireinject

package main

import (
	"W-chat/config"
	"W-chat/src/httpserver"
	"W-chat/src/methods"
	"W-chat/src/repository"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	repository.NewMySQLClient,
	methods.ProviderSet,
)

func HttpServerInjector(conf *config.Config) *httpserver.Basic {
	panic(
		wire.Build(
			providerSet,
			httpserver.ProviderSet,
		),
	)
}
