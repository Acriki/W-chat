package httpserver

import (
	"W-chat/src/httpserver/api/handler"
	"W-chat/src/httpserver/api/router"
	"W-chat/src/repository/cache"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler.ProviderSet,
	cache.ProviderSet,
	router.NewRouter,
	wire.Struct(new(Basic), "*"),
)
