package routes

import (
	"ects/internal/controllers/auth"
	"github.com/kataras/iris/mvc"
)

func authentication(application *mvc.Application) {
	application.Handle(new(auth.AuthenticationController))
}