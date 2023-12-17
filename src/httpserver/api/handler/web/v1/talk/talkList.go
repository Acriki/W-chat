package talk

import (
	"W-chat/message/pb/web/v1"
	tf "W-chat/pkg/time"
	mygin "W-chat/src/gin"
	"W-chat/src/methods/talk"
	"W-chat/src/repository/cache"
	"W-chat/src/repository/database"
	"fmt"
)

type Session struct {
	ContactRepo *database.Contact
	UsersRepo   *database.Users
	GroupRepo   *database.Group

	MessageCache *cache.MessageCache
	UnreadCache  *cache.UnreadCache
	ClientCache  *cache.ClientCache

	TalkSession *talk.TalkSessionMethods
}

// List 会话列表
func (c *Session) List(ctx *mygin.Context) error {

	uid := ctx.UserId()

	// 获取未读消息数
	unReads := c.UnreadCache.All(ctx.Ctx(), uid)
	if len(unReads) > 0 {
		c.TalkSession.BatchAddList(ctx.Ctx(), uid, unReads)
	}

	data, err := c.TalkSession.List(ctx.Ctx(), uid)
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	friends := make([]int, 0)
	for _, item := range data {
		if item.TalkType == 1 {
			friends = append(friends, item.ReceiverId)
		}
	}

	// 获取好友备注
	remarks, _ := c.ContactRepo.Remarks(ctx.Ctx(), uid, friends)

	items := make([]*web.TalkSessionItem, 0)
	for _, item := range data {
		value := &web.TalkSessionItem{
			Id:         int32(item.Id),
			TalkType:   int32(item.TalkType),
			ReceiverId: int32(item.ReceiverId),
			IsTop:      int32(item.IsTop),
			IsDisturb:  int32(item.IsDisturb),
			IsRobot:    int32(item.IsRobot),
			Avatar:     item.UserAvatar,
			MsgText:    "...",
			UpdatedAt:  tf.FormatDatetime(item.UpdatedAt),
		}

		if num, ok := unReads[fmt.Sprintf("%d_%d", item.TalkType, item.ReceiverId)]; ok {
			value.UnreadNum = int32(num)
		}

		if item.TalkType == 1 {
			value.Name = item.Nickname
			value.Avatar = item.UserAvatar
			value.Remark = remarks[item.ReceiverId]
			// value.IsOnline = int32(strutil.BoolToInt(c.ClientCache.IsOnline(ctx.Ctx(), entity.ImChannelChat, strconv.Itoa(int(value.ReceiverId)))))
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		// 查询缓存消息
		if msg, err := c.MessageCache.Get(ctx.Ctx(), item.TalkType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}

		items = append(items, value)
	}

	return ctx.Success(&web.TalkSessionListResponse{Items: items})
}
