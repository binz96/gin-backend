package v1

import (
	"net/http"
	"strconv"

	"github.com/binz96/blog/errmsg"
	"github.com/binz96/blog/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var u model.User
	_ = c.ShouldBindJSON(&u)

	code := model.CheckUser(u.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&u)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    u,
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	users := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    users,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u model.User
	c.ShouldBindJSON(&u)
	code := model.CheckUser(u.Username)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &u)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}
