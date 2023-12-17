package user

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Struct(new(Account), "*"),
	wire.Struct(new(User), "*"),
	wire.Struct(new(Auth), "*"),
)
