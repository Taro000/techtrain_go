package server

import (
	"github.com/gin-gonic/gin"
	"techtrain_go_api/controller"
)

// Initialize server
func Init()  {
	r := router()
	r.Run()
}


// Routing
func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/user")
	{
		ctrl := controller.UserController{}
		u.GET("/get", ctrl.GetOneName)
		u.POST("/create", ctrl.Create)
		u.PUT("/update", ctrl.Update)
	}

	uc := r.Group("/gacha")
	{
		ctrl := controller.UserCharacterController{}
		uc.POST("/draw", ctrl.Gacha)
	}

	c := r.Group("/character")
	{
		ctrl := controller.CharacterController{}
		c.GET("/list", ctrl.GetCharacters)
		c.POST("/create", ctrl.Create)
		c.PUT("update", ctrl.Update)
	}

	return r
}