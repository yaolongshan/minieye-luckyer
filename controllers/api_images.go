package controllers

import (
	"code/minieye-luckyer/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ApiImagesPreview(c *gin.Context) {
	imageName := c.Param("name")
	path := fmt.Sprintf("%v/%v/%v", conf.Conf.RootPath, "images", imageName)
	//强制缓存
	c.Header("Cache-Control", "private, max-age=86400")
	c.File(path)
}
