package web

import v1 "W-chat/src/httpserver/api/handler/web/v1"

type V1 struct {
	Auth *v1.Auth
}

type V2 struct {
}

type Handler struct {
	V1 *V1
}
