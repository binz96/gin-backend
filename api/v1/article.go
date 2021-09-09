package v1

import (
	"net/http"
	"strconv"

	"github.com/binz96/blog/errmsg"
	"github.com/binz96/blog/model"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var a model.Article
	_ = c.ShouldBindJSON(&a)

	code := model.CreateArticle(&a)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    a,
	})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	a, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    a,
	})
}

func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	articles, code := model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"data":    articles,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var a model.Article
	c.ShouldBindJSON(&a)

	code := model.UpdateArticle(id, &a)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
	})
}
