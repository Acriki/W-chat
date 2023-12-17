package cache

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewClientStorage,
	NewMessageStorage,
	NewTokenSessionCache,
	NewContactRemark,
	NewRelation,
	NewUnreadCache,
)
