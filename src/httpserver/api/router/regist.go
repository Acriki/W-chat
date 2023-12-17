package router

import (
	"W-chat/config"
	mygin "W-chat/src/gin"
	"W-chat/src/gin/middleware"
	"W-chat/src/httpserver/api/handler/web"
	"W-chat/src/repository/cache"

	"github.com/gin-gonic/gin"
)

// RegisterWebRoute 注册 Web 路由
func RegisterWebRoute(router *gin.Engine, handler *web.Handler, jwtCache *cache.JwtTokenCache, conf *config.Config) {

	// 授权验证中间件
	authorize := middleware.Auth(conf.Jwt.Secret, "api", jwtCache)

	// v1 接口
	v1 := router.Group("/api/v1")
	{
		// 授权相关分组
		auth := v1.Group("/auth")
		{
			auth.POST("/login", mygin.HandlerFunc(handler.V1.User.Auth.Login)) // 登录
			// auth.POST("/register", mygin.HandlerFunc(handler.V1.Auth.Register))          // 注册
			// auth.POST("/refresh", authorize, mygin.HandlerFunc(handler.V1.Auth.Refresh)) // 刷新 Token
			// auth.POST("/logout", authorize, mygin.HandlerFunc(handler.V1.Auth.Logout))   // 退出登录
			// auth.POST("/forget", mygin.HandlerFunc(handler.V1.Auth.Forget))              // 找回密码
		}

		// 用户相关分组
		user := v1.Group("/users").Use(authorize)
		{
			user.GET("/detail", mygin.HandlerFunc(handler.V1.User.Account.Detail))   // 获取个人信息
			user.GET("/setting", mygin.HandlerFunc(handler.V1.User.Account.Setting)) // 获取个人信息
			// user.POST("/change/detail", mygin.HandlerFunc(handler.V1.User.ChangeDetail))     // 修改用户信息
			// user.POST("/change/password", mygin.HandlerFunc(handler.V1.User.ChangePassword)) // 修改用户密码
			// user.POST("/change/mobile", mygin.HandlerFunc(handler.V1.User.ChangeMobile))     // 修改用户手机号
			// user.POST("/change/email", mygin.HandlerFunc(handler.V1.User.ChangeEmail))       // 修改用户邮箱
		}

		talk := v1.Group("/talk").Use(authorize)
		{
			talk.GET("/list", mygin.HandlerFunc(handler.V1.Talk.TalkList.List)) // 会话列表
			// talk.POST("/create", mygin.HandlerFunc(handler.V1.Talk.Create))                              // 创建会话
			// talk.POST("/delete", mygin.HandlerFunc(handler.V1.Talk.Delete))                              // 删除会话
			// talk.POST("/topping", mygin.HandlerFunc(handler.V1.Talk.Top))                                // 置顶会话
			// talk.POST("/disturb", mygin.HandlerFunc(handler.V1.Talk.Disturb))                            // 会话免打扰
			// talk.GET("/records", mygin.HandlerFunc(handler.V1.TalkRecords.GetRecords))                   // 会话面板记录
			// talk.GET("/records/history", mygin.HandlerFunc(handler.V1.TalkRecords.SearchHistoryRecords)) // 历史会话记录
			// talk.GET("/records/forward", mygin.HandlerFunc(handler.V1.TalkRecords.GetForwardRecords))    // 会话转发记录
			// talk.GET("/records/file/download", mygin.HandlerFunc(handler.V1.TalkRecords.Download))       // 会话转发记录
			// talk.POST("/unread/clear", mygin.HandlerFunc(handler.V1.Talk.ClearUnreadMessage))            // 清除会话未读数
		}

		// contact := v1.Group("/contact").Use(authorize)
		{
			// contact.GET("/list", mygin.HandlerFunc(handler.V1.Contact.List))             // 联系人列表
			// contact.GET("/search", mygin.HandlerFunc(handler.V1.Contact.Search))         // 搜索联系人
			// contact.GET("/detail", mygin.HandlerFunc(handler.V1.Contact.Detail))         // 搜索联系人
			// contact.POST("/delete", mygin.HandlerFunc(handler.V1.Contact.Delete))        // 删除联系人
			// contact.POST("/edit-remark", mygin.HandlerFunc(handler.V1.Contact.Remark))   // 编辑联系人备注
			// contact.POST("/move-group", mygin.HandlerFunc(handler.V1.Contact.MoveGroup)) // 编辑联系人备注

			// // 联系人申请相关
			// contact.GET("/apply/records", mygin.HandlerFunc(handler.V1.ContactApply.List))              // 联系人申请列表
			// contact.POST("/apply/create", mygin.HandlerFunc(handler.V1.ContactApply.Create))            // 添加联系人申请
			// contact.POST("/apply/accept", mygin.HandlerFunc(handler.V1.ContactApply.Accept))            // 同意人申请列表
			// contact.POST("/apply/decline", mygin.HandlerFunc(handler.V1.ContactApply.Decline))          // 拒绝人申请列表
			// contact.GET("/apply/unread-num", mygin.HandlerFunc(handler.V1.ContactApply.ApplyUnreadNum)) // 联系人申请未读数

			// 联系人分组
			// contact.GET("/group/list", mygin.HandlerFunc(handler.V1.ContactGroup.List))  // 联系人分组列表
			// contact.POST("/group/save", mygin.HandlerFunc(handler.V1.ContactGroup.Save)) // 联系人分组排序
		}

		// 聊天群相关分组
		// userGroup := v1.Group("/group").Use(authorize)
		{
			// userGroup.GET("/list", mygin.HandlerFunc(handler.V1.Group.GroupList))            // 群组列表
			// userGroup.GET("/overt/list", mygin.HandlerFunc(handler.V1.Group.OvertList))      // 公开群组列表
			// userGroup.GET("/detail", mygin.HandlerFunc(handler.V1.Group.Detail))             // 群组详情
			// userGroup.POST("/create", mygin.HandlerFunc(handler.V1.Group.Create))            // 创建群组
			// userGroup.POST("/dismiss", mygin.HandlerFunc(handler.V1.Group.Dismiss))          // 解散群组
			// userGroup.POST("/invite", mygin.HandlerFunc(handler.V1.Group.Invite))            // 邀请加入群组
			// userGroup.POST("/secede", mygin.HandlerFunc(handler.V1.Group.SignOut))           // 退出群组
			// userGroup.POST("/setting", mygin.HandlerFunc(handler.V1.Group.Setting))          // 设置群组信息
			// userGroup.POST("/handover", mygin.HandlerFunc(handler.V1.Group.Handover))        // 群主转让
			// userGroup.POST("/assign-admin", mygin.HandlerFunc(handler.V1.Group.AssignAdmin)) // 分配管理员
			// userGroup.POST("/no-speak", mygin.HandlerFunc(handler.V1.Group.NoSpeak))         // 修改禁言状态
			// userGroup.POST("/mute", mygin.HandlerFunc(handler.V1.Group.Mute))                // 修改禁言状态
			// userGroup.POST("/overt", mygin.HandlerFunc(handler.V1.Group.Overt))              // 修改禁言状态

			// // 群成员相关
			// userGroup.GET("/member/list", mygin.HandlerFunc(handler.V1.Group.Members))               // 群成员列表
			// userGroup.GET("/member/invites", mygin.HandlerFunc(handler.V1.Group.GetInviteFriends))   // 群成员列表
			// userGroup.POST("/member/remove", mygin.HandlerFunc(handler.V1.Group.RemoveMembers))      // 移出指定群成员
			// userGroup.POST("/member/remark", mygin.HandlerFunc(handler.V1.Group.UpdateMemberRemark)) // 设置群名片

			// // 群公告相关
			// userGroup.GET("/notice/list", mygin.HandlerFunc(handler.V1.GroupNotice.List))             // 群公告列表
			// userGroup.POST("/notice/edit", mygin.HandlerFunc(handler.V1.GroupNotice.CreateAndUpdate)) // 添加或编辑群公告
			// userGroup.POST("/notice/delete", mygin.HandlerFunc(handler.V1.GroupNotice.Delete))        // 删除群公告

			// // 群申请
			// userGroup.POST("/apply/create", mygin.HandlerFunc(handler.V1.GroupApply.Create))        // 提交入群申请
			// userGroup.POST("/apply/agree", mygin.HandlerFunc(handler.V1.GroupApply.Agree))          // 同意入群申请
			// userGroup.POST("/apply/decline", mygin.HandlerFunc(handler.V1.GroupApply.Decline))      // 拒绝入群申请
			// userGroup.GET("/apply/list", mygin.HandlerFunc(handler.V1.GroupApply.List))             // 入群申请列表
			// userGroup.GET("/apply/all", mygin.HandlerFunc(handler.V1.GroupApply.All))               // 入群申请列表
			// userGroup.GET("/apply/unread", mygin.HandlerFunc(handler.V1.GroupApply.ApplyUnreadNum)) // 入群申请未读
		}

	}

}
