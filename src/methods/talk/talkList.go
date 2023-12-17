package talk

import (
	tf "W-chat/pkg/time"
	"W-chat/src/repository/database"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TalkSessionMethods struct {
	cache       *redis.Client
	db          *gorm.DB
	talkSession *database.TalkSession
}

func NewTalkSessionMethodsObj(cache *redis.Client, db *gorm.DB, talkSession *database.TalkSession) *TalkSessionMethods {
	return &TalkSessionMethods{cache, db, talkSession}
}

func (s *TalkSessionMethods) List(ctx context.Context, uid int) ([]*database.SearchTalkSession, error) {

	fields := []string{
		"list.id", "list.talk_type", "list.receiver_id", "list.updated_at",
		"list.is_disturb", "list.is_top", "list.is_robot",
		"`users`.avatar as user_avatar", "`users`.nickname",
		"`group`.group_name", "`group`.avatar as group_avatar",
	}

	query := s.db.WithContext(ctx).Table("talk_session list")
	query.Joins("left join `users` ON list.receiver_id = `users`.id AND list.talk_type = 1")
	query.Joins("left join `group` ON list.receiver_id = `group`.id AND list.talk_type = 2")
	query.Where("list.user_id = ? and list.is_delete = 0", uid)
	query.Order("list.updated_at desc")

	var items []*database.SearchTalkSession
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type TalkSessionCreateOpt struct {
	UserId     int
	TalkType   int
	ReceiverId int
	IsBoot     bool
}

// Create 创建会话列表
func (s *TalkSessionMethods) Create(ctx context.Context, opt *TalkSessionCreateOpt) (*database.TalkSessionModel, error) {

	result, err := s.talkSession.FindByWhere(ctx, "talk_type = ? and user_id = ? and receiver_id = ?", opt.TalkType, opt.UserId, opt.ReceiverId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = &database.TalkSessionModel{
			TalkType:   opt.TalkType,
			UserId:     opt.UserId,
			ReceiverId: opt.ReceiverId,
		}

		if opt.IsBoot {
			result.IsRobot = 1
		}

		s.db.WithContext(ctx).Create(result)
	} else {
		result.IsTop = 0
		result.IsDelete = 0
		result.IsDisturb = 0

		if opt.IsBoot {
			result.IsRobot = 1
		}

		s.db.WithContext(ctx).Save(result)
	}

	return result, nil
}

// Delete 删除会话
func (s *TalkSessionMethods) Delete(ctx context.Context, uid int, id int) error {
	_, err := s.talkSession.UpdateWhere(ctx, map[string]any{"is_delete": 1, "updated_at": time.Now()}, "id = ? and user_id = ?", id, uid)
	return err
}

type TalkSessionTopOpt struct {
	UserId int
	Id     int
	Type   int
}

// Top 会话置顶
func (s *TalkSessionMethods) Top(ctx context.Context, opt *TalkSessionTopOpt) error {
	_, err := s.talkSession.UpdateWhere(ctx, map[string]any{
		"is_top":     opt.Type,
		"updated_at": time.Now(),
	}, "id = ? and user_id = ?", opt.Id, opt.UserId)
	return err
}

type TalkSessionDisturbOpt struct {
	UserId     int
	TalkType   int
	ReceiverId int
	IsDisturb  int
}

// Disturb 会话免打扰
func (s *TalkSessionMethods) Disturb(ctx context.Context, opt *TalkSessionDisturbOpt) error {
	_, err := s.talkSession.UpdateWhere(ctx, map[string]any{
		"is_disturb": opt.IsDisturb,
		"updated_at": time.Now(),
	}, "user_id = ? and receiver_id = ? and talk_type = ?", opt.UserId, opt.ReceiverId, opt.TalkType)
	return err
}

// BatchAddList 批量添加会话列表
func (s *TalkSessionMethods) BatchAddList(ctx context.Context, uid int, values map[string]int) {

	ctime := tf.DateTime()

	data := make([]string, 0)
	for k, v := range values {
		if v == 0 {
			continue
		}

		value := strings.Split(k, "_")
		if len(value) != 2 {
			continue
		}

		data = append(data, fmt.Sprintf("(%s, %d, %s, '%s', '%s')", value[0], uid, value[1], ctime, ctime))
	}

	if len(data) == 0 {
		return
	}

	s.db.WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO talk_session ( `talk_type`, `user_id`, `receiver_id`, created_at, updated_at ) VALUES %s ON DUPLICATE KEY UPDATE is_delete = 0, updated_at = '%s'", strings.Join(data, ","), ctime))
}
