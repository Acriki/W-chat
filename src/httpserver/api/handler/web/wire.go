package web

import (
	"W-chat/src/httpserver/api/handler/web/v1/talk"
	"W-chat/src/httpserver/api/handler/web/v1/user"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.ProviderSet,
	talk.ProviderSet,
	wire.Struct(new(V1), "*"),
	wire.Struct(new(Handler), "*"),
)
