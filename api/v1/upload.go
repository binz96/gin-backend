package v1

import (
	"net/http"

	"github.com/binz96/blog/errmsg"
	"github.com/binz96/blog/model"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.CodeToMsg(code),
		"url":     url,
	})
}
