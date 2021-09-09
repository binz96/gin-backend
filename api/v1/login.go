package v1

import (
	"net/http"

	"github.com/binz96/blog/errmsg"
	"github.com/binz96/blog/middleware"
	"github.com/binz96/blog/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.User
	c.ShouldBindJSON(&u)
	var token string
	var code int
	code = model.CheckLogin(u.Username, u.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(u.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"token":   token,
	})
}
