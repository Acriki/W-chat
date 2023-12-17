package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UnreadCache struct {
	redis *redis.Client
}

func NewUnreadCache(rds *redis.Client) *UnreadCache {
	return &UnreadCache{rds}
}

// Incr 消息未读数自增
// @params mode    对话模式 1私信 2群聊
// @params sender  发送者ID
// @params receive 接收者ID
func (u *UnreadCache) Incr(ctx context.Context, mode, sender, receive int) {
	u.redis.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

// PipeIncr 消息未读数自增
// @params mode    对话模式 1私信 2群聊
// @params sender  发送者ID
// @params receive 接收者ID
func (u *UnreadCache) PipeIncr(ctx context.Context, pipe redis.Pipeliner, mode, sender, receive int) {
	pipe.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

// Get 获取消息未读数
// @params mode    对话模式 1私信 2群聊
// @params sender  发送者ID
// @params receive 接收者ID
func (u *UnreadCache) Get(ctx context.Context, mode, sender, receive int) int {
	val, _ := u.redis.HGet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender)).Int()
	return val
}

// Del 删除消息未读数
// @params mode    对话模式 1私信 2群聊
// @params sender  发送者ID
// @params receive 接收者ID
func (u *UnreadCache) Del(ctx context.Context, mode, sender, receive int) {
	u.redis.HDel(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender))
}

// Reset 消息未读数重置
// @params mode    对话模式 1私信 2群聊
// @params sender  发送者ID
// @params receive 接收者ID
func (u *UnreadCache) Reset(ctx context.Context, mode, sender, receive int) {
	u.redis.HSet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 0)
}

func (u *UnreadCache) All(ctx context.Context, receive int) map[string]int {
	items := make(map[string]int)
	for k, v := range u.redis.HGetAll(ctx, u.name(receive)).Val() {
		items[k], _ = strconv.Atoi(v)
	}

	return items
}

func (u *UnreadCache) name(receive int) string {
	return fmt.Sprintf("im:message:unread:uid_%d", receive)
}
