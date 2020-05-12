package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"techtrain_go_api/db"
	"techtrain_go_api/models"
	"time"
)

// Provides Character's behavior
type CharacterService struct {}

//Alias of models/Character
type Character models.Character

/*
Get list of Character which a User has
param: User.token
*/
func (CS CharacterService) CharactersHadByUser(token string) ([]Character, error) {
	sqlQuery := "SELECT user_characters.user_character_id, characters.id, characters.name " +
		        "FROM users " +
		        "JOIN user_characters ON user_characters.user_id = users.id " +
		        "JOIN characters ON characters.id = user_characters.character_id " +
		        "AND users.token = ?"
	db := db.GetDB()
	var characters []Character

	if err := db.Raw(sqlQuery, token).Find(&characters).Error; err != nil{
		return nil, err
	}

	return characters, nil
}


// Get all Characters
func (CS CharacterService) GetAll() ([]Character, error) {
	db := db.GetDB()
	var characters []Character

	if err := db.Find(&characters).Error; err != nil {
		return nil, err
	}

	return characters, nil
}


//Create Character model
func (CS CharacterService) Insert(context *gin.Context) (Character, error) {
	db := db.GetDB()
	var character Character

	uu, err := GenerateUUID()
	if err != nil{
		return character, err
	}
	character.ID = uu
	character.Name = context.PostForm("name")
	probability, _ := strconv.Atoi(context.PostForm("probability"))
	character.Probability = probability
	character.CreatedAt = time.Now()
	character.ModifiedAt = time.Now()

	if err := context.BindJSON(&character); err != nil {
		return character, err
	}

	if err := db.Create(&character).Error; err != nil {
		return character, err
	}

	return character, nil
}


//Get a Character
func (CS CharacterService) GetOne(id string) (Character, error) {
	db := db.GetDB()
	var character Character

	if err := db.Where("id = ?", id).First(&character).Error; err != nil {
		return character, err
	}

	return character, nil
}


//Update a Character
func (CS CharacterService) Update(id string, context *gin.Context) (Character, error) {
	db := db.GetDB()
	var character Character

	if err := db.Where("id = ?", id).First(&character).Error; err != nil {
		return character, err
	}

	character.Name = context.PostForm("name")
	probability, _ := strconv.Atoi(context.PostForm("probability"))
	character.Probability = probability
	character.ModifiedAt = time.Now()

	if err := context.BindJSON(&character); err != nil {
		return character, err
	}

	db.Save(&character)

	return character, nil
}


//Delete a Character
func (CS CharacterService) Delete(id string, context *gin.Context) error {
	db := db.GetDB()
	var character Character

	if err := db.Where("id = ?", id).Delete(&character).Error; err != nil {
		return err
	}

	return nil
}