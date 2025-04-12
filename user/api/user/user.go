package user

import (
	"github.com/gin-gonic/gin"
	"user/pkg/repo"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(r *gin.Engine) {}
