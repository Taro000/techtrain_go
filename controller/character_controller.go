package controller


import (
	"fmt"
	"github.com/gin-gonic/gin"

	"techtrain_go_api/service"
)


//Controller of User
type CharacterController struct {}


func (CC CharacterController) GetCharacters(context *gin.Context) {
	token := context.Request.Header.Get("x-token")
	var cs service.CharacterService

	characters, err := cs.CharactersHadByUser(token)
	if err != nil{
		context.AbortWithStatus(404)
		fmt.Println(err)
	}

	context.JSON(200, characters)
}

func (CC CharacterController) Create(context *gin.Context) {
	var cs service.CharacterService

	character, err := cs.Insert(context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	}else {
		context.JSON(201, character.Name)
	}
}


func (CC CharacterController) Update(context *gin.Context) {
	id := context.Request.Header.Get("character-id")
	var cs service.CharacterService

	_, err := cs.Update(id, context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	}else {
		context.JSON(200, nil)
	}
}