package web

import (
	"W-chat/src/httpserver/api/handler/web/v1/talk"
	"W-chat/src/httpserver/api/handler/web/v1/user"
)

type V1 struct {
	User *user.User
	Talk *talk.Talk
}

type V2 struct {
}

type Handler struct {
	V1 *V1
}
