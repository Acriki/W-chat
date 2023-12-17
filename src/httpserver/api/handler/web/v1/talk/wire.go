package talk

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Struct(new(Session), "*"),
	wire.Struct(new(Talk), "*"),
)
