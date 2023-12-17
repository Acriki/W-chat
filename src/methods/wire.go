package methods

import (
	"W-chat/src/methods/talk"
	"W-chat/src/methods/user"
	"W-chat/src/repository/database"

	"github.com/google/wire"
)

type Methods struct {
	Auth     *user.UserAuthMethods
	TalkList *talk.TalkSessionMethods
}

var ProviderSet = wire.NewSet(
	database.ProviderSet,
	user.NewUserAuthMethodsObj,
	talk.NewTalkSessionMethodsObj,
	wire.Struct(new(Methods), "*"),
)
