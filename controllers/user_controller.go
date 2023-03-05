package controllers

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	c.String(200, "Register ..")
}

func Login(c *gin.Context) {
	c.String(200, "Login ..")
}
