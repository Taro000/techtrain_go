package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"techtrain_go_api/service"
)


//Controller of User
type UserController struct {}


func (UC UserController) GetOneName(context *gin.Context) {
	token := context.Request.Header.Get("x-token")
	var us service.UserService

	user, err := us.GetOne(token)

	if err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	}else {
		context.JSON(200, user.Name)
	}
}


func (UC UserController) Create(context *gin.Context) {
	var us service.UserService

	user, err := us.Insert(context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	}else {
		context.JSON(201, user.Token)
	}
}


func (UC UserController) Update(context *gin.Context) {
	//copied_context := context.Copy()
	token := context.Request.Header.Get("x-token")
	var us service.UserService

	//bodyCopy := new(bytes.Buffer)
	//_, err := io.Copy(bodyCopy, context.Request.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//bodyData := bodyCopy.Bytes()
	//context.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
	//fmt.Println(context.Request.Body)

	_, err := us.Update(token, context)

	if err != nil {
		context.AbortWithStatus(400)
		fmt.Println(err)
	}else {
		context.JSON(200, nil)
	}
}