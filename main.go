package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/registration", func(c *gin.Context) {
		var content Registr
		session := sessions.Default(c)
		if err := c.ShouldBindJSON(&content); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		answer := Registration(content, session)
		c.JSON(200, answer)
	})

	router.POST("/confirmation", func(c *gin.Context) {
		session := sessions.Default(c)
		var userCode Сonfirm
		if err := c.ShouldBindJSON(&userCode); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		answer := Confirmation(userCode, session)

		c.JSON(200, answer)

		//c.JSON(200, answer)
	})

	router.POST("/authorization", func(c *gin.Context) {
		session := sessions.Default(c)
		var user Authorization

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		answer := Authorizations(user, session)
		c.JSON(200, answer)

	})

	router.POST("/settings", func(c *gin.Context) {
		sessions := sessions.Default(c)
		var setting Setting
		if err := c.ShouldBindJSON(&setting); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		answer := Settings(setting, sessions)
		c.JSON(200, answer)
	})

	router.Run("0.0.0.0:8080")
}
