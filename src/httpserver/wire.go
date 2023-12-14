package httpserver

import (
	"W-chat/src/httpserver/api/handler"
	"W-chat/src/httpserver/api/router"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler.ProviderSet,
	router.NewRouter,
	wire.Struct(new(Basic), "*"),
)
