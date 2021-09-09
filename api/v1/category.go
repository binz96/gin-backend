package v1

import (
	"net/http"
	"strconv"

	"github.com/binz96/blog/errmsg"
	"github.com/binz96/blog/model"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var ca model.Category
	_ = c.ShouldBindJSON(&ca)

	code := model.CheckCategory(ca.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&ca)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    ca,
	})
}

func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	cates := model.GetCategories(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    cates,
	})
}

func GetArticlesInCategory(c *gin.Context) {
	cid := c.Param("cid")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	articles, code := model.GetArticlesInCategory(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    articles,
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var ca model.Category
	c.ShouldBindJSON(&ca)
	code := model.CheckCategory(ca.Name)
	if code == errmsg.SUCCESS {
		model.UpdateCategory(id, &ca)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}
