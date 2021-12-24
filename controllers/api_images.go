package controllers

import (
	"code/minieye-luckyer/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ApiImages(c *gin.Context) {
	imageName := c.Param("name")
	path := fmt.Sprintf("%v/%v/%v", conf.Conf.RootPath, "images", imageName)
	c.File(path)
}
