package service

import (
	"github.com/gin-gonic/gin"
	"techtrain_go_api/db"
	"techtrain_go_api/models"
	"time"
)


// Provides User's behavior
type UserService struct {}

//Alias of models/User
type User models.User


//Create User model
func (US UserService) Insert(context *gin.Context) (User, error) {
	db := db.GetDB()
	var user User

	uu1, err := GenerateUUID()
	if err != nil{
		return user, err
	}
	uu2, err := GenerateUUID()
	if err != nil{
		return user, err
	}
	user.ID = uu1
	user.Token = uu2
	user.Name = context.PostForm("name")
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()

	if err := context.BindJSON(&user); err != nil {
		return user, err
	}

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}


//Get a User(By Token)
func (US UserService) GetOne(token string) (User, error) {
	db := db.GetDB()
	var user User

	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}


//Update a User
func (US UserService) Update(token string, context *gin.Context) (User, error) {
	db := db.GetDB()
	var user User
	//var json_map map[string]string

	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		return user, err
	}

	//buf := make([]byte, 2048)
	//n, _ := context.Request.Body.Read(buf)
	//b := buf[0:n]
	//
	//if err := json.Unmarshal(b, &json_map); err != nil{
	//	return user, err
	//}
	//
	//user.Name = json_map["name"]
	user.ModifiedAt = time.Now()

	if err := context.BindJSON(&user); err != nil {
		return user, err
	}

	db.Save(&user)

	return user, nil
}


//Delete a User
func (US UserService) Delete(token string, context *gin.Context) error {
	db := db.GetDB()
	var user User

	if err := db.Where("token = ?", token).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}