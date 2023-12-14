package handler

import (
	"W-chat/src/httpserver/api/handler/web"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.ProviderSet,
	wire.Struct(new(Handler), "*"),
)
