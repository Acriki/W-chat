// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"W-chat/config"
	"W-chat/src/httpserver"
	"W-chat/src/httpserver/api/handler"
	"W-chat/src/httpserver/api/handler/web"
	talk2 "W-chat/src/httpserver/api/handler/web/v1/talk"
	user2 "W-chat/src/httpserver/api/handler/web/v1/user"
	"W-chat/src/httpserver/api/router"
	"W-chat/src/methods"
	"W-chat/src/methods/talk"
	"W-chat/src/methods/user"
	"W-chat/src/repository"
	"W-chat/src/repository/cache"
	"W-chat/src/repository/database"
	"github.com/google/wire"
)

// Injectors from wire.go:

func HttpServerInjector(conf *config.Config) *httpserver.Basic {
	db := repository.NewMySQLClient(conf)
	users := database.NewUsers(db)
	userAuthMethods := user.NewUserAuthMethodsObj(users)
	client := repository.NewRedisClient(conf)
	talkSession := database.NewTalkSession(db)
	talkSessionMethods := talk.NewTalkSessionMethodsObj(client, db, talkSession)
	methodsMethods := &methods.Methods{
		Auth:     userAuthMethods,
		TalkList: talkSessionMethods,
	}
	auth := &user2.Auth{
		Config:  conf,
		Methods: methodsMethods,
	}
	account := &user2.Account{
		UsersRepo: users,
	}
	userUser := &user2.User{
		Auth:    auth,
		Account: account,
	}
	contactRemark := cache.NewContactRemark(client)
	relation := cache.NewRelation(client)
	contact := database.NewContact(db, contactRemark, relation)
	group := database.NewGroup(db)
	messageCache := cache.NewMessageStorage(client)
	unreadCache := cache.NewUnreadCache(client)
	clientCache := cache.NewClientStorage(client, conf)
	session := &talk2.Session{
		ContactRepo:  contact,
		UsersRepo:    users,
		GroupRepo:    group,
		MessageCache: messageCache,
		UnreadCache:  unreadCache,
		ClientCache:  clientCache,
		TalkSession:  talkSessionMethods,
	}
	talkTalk := &talk2.Talk{
		TalkList: session,
	}
	v1 := &web.V1{
		User: userUser,
		Talk: talkTalk,
	}
	webHandler := &web.Handler{
		V1: v1,
	}
	handlerHandler := &handler.Handler{
		WebApi: webHandler,
	}
	jwtTokenCache := cache.NewTokenSessionCache(client)
	engine := router.NewRouter(conf, handlerHandler, jwtTokenCache)
	basic := &httpserver.Basic{
		Config: conf,
		Gin:    engine,
	}
	return basic
}

// wire.go:

var providerSet = wire.NewSet(repository.NewMySQLClient, repository.NewRedisClient, methods.ProviderSet)
