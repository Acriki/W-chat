package database

import (
	"context"
	"strconv"
	"time"

	"W-chat/src/repository/cache"

	"gorm.io/gorm"
)

const (
	ContactStatusNormal = 1
	ContactStatusDelete = 0
)

// ContactModel 用户好友关系表
type ContactModel struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                         // 关系ID
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`                       // 用户id
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL" json:"friend_id"`                   // 好友id
	Remark    string    `gorm:"column:remark;NOT NULL" json:"remark"`                                   // 好友的备注
	Status    int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                         // 好友状态 [0:否;1:是]
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`                     // 分组id
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"` // 更新时间
}

func (ContactModel) TableName() string {
	return "contact"
}

type ContactListItem struct {
	Id       int    `gorm:"column:id" json:"id"`                // 用户ID
	Nickname string `gorm:"column:nickname" json:"nickname"`    // 用户昵称
	Gender   uint8  `gorm:"column:gender" json:"gender"`        // 用户性别[0:未知;1:男;2:女;]
	Motto    string `gorm:"column:motto" json:"motto"`          // 用户座右铭
	Avatar   string `grom:"column:avatar" json:"avatar" `       // 好友头像
	Remark   string `gorm:"column:remark" json:"friend_remark"` // 好友的备注
	IsOnline int    `json:"is_online"`                          // 是否在线
	GroupId  int    `gorm:"column:group_id" json:"group_id"`    // 联系人分组
}

type Contact struct {
	Repo[ContactModel]
	remarkCache *cache.ContactRemark
	relation    *cache.Relation
}

func NewContact(db *gorm.DB, remarkCache *cache.ContactRemark, relation *cache.Relation) *Contact {
	return &Contact{Repo: NewRepo[ContactModel](db), remarkCache: remarkCache, relation: relation}
}

func (c *Contact) Remarks(ctx context.Context, uid int, fids []int) (map[int]string, error) {

	if !c.remarkCache.Exist(ctx, uid) {
		_ = c.LoadContactCache(ctx, uid)
	}

	return c.remarkCache.MGet(ctx, uid, fids)
}

// IsFriend 判断是否为好友关系
func (c *Contact) IsFriend(ctx context.Context, uid int, friendId int, remarkCache bool) bool {

	if remarkCache && c.relation.IsContactRelation(ctx, uid, friendId) == nil {
		return true
	}

	count, err := c.Repo.QueryCount(ctx, "((user_id = ? and friend_id = ?) or (user_id = ? and friend_id = ?)) and status = ?", uid, friendId, friendId, uid, ContactStatusNormal)
	if err != nil {
		return false
	}

	if count == 2 {
		c.relation.SetContactRelation(ctx, uid, friendId)
	} else {
		c.relation.DelContactRelation(ctx, uid, friendId)
	}

	return count == 2
}

func (c *Contact) GetFriendRemark(ctx context.Context, uid int, friendId int) string {

	if c.remarkCache.Exist(ctx, uid) {
		return c.remarkCache.Get(ctx, uid, friendId)
	}

	var remark string
	c.Repo.Model(ctx).Where("user_id = ? and friend_id = ?", uid, friendId).Pluck("remark", &remark)

	return remark
}

func (c *Contact) SetFriendRemark(ctx context.Context, uid int, friendId int, remark string) error {
	return c.remarkCache.Set(ctx, uid, friendId, remark)
}

func (c *Contact) LoadContactCache(ctx context.Context, uid int) error {

	all, err := c.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Select("friend_id,remark").Where("user_id = ? and status = ?", uid, ContactStatusNormal)
	})

	if err != nil {
		return err
	}

	items := make(map[string]any)
	for _, value := range all {
		if len(value.Remark) > 0 {
			items[strconv.Itoa(value.FriendId)] = value.Remark
		}
	}

	return c.remarkCache.MSet(ctx, uid, items)
}
