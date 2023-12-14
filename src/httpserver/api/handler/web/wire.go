package web

import (
	v1 "W-chat/src/httpserver/api/handler/web/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(v1.Auth), "*"),
	wire.Struct(new(V1), "*"),
	wire.Struct(new(Handler), "*"),
)