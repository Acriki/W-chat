package methods

import (
	"W-chat/src/repository/database"

	"github.com/google/wire"
)

type Methods struct {
	Auth *AuthMethods
}

var ProviderSet = wire.NewSet(
	database.ProviderSet,
	NewAuthMethodsObj,
	wire.Struct(new(Methods), "*"),
)