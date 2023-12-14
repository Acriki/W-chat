//go:build wireinject
// +build wireinject

package main

import (
	"W-chat/config"
	"W-chat/src/httpserver"

	"github.com/google/wire"
)

// // var providerSet = wire.NewSet(
// // 	database.ProviderSet,
// // 	repository.NewMySQLClient,
// // )

func HttpServerInjector(conf *config.Config) *httpserver.Basic {
	panic(
		wire.Build(
			// database.ProviderSet,
			// repository.NewMySQLClient,
			httpserver.ProviderSet,
		),
	)
}
