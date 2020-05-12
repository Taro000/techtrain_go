package service

import (
	"techtrain_go_api/db"
	"techtrain_go_api/models"
	"time"
)


// Provides UserCharacter's behavior
type UserCharacterService struct {}

//Alias of models/UserCharacter
type UserCharacter models.UserCharacter


// Get all UserCharacters
func (UCS UserCharacterService) GetAll() ([]UserCharacter, error) {
	db := db.GetDB()
	var userCharacterList []UserCharacter

	if err := db.Find(&userCharacterList).Error; err != nil {
		return nil, err
	}

	return userCharacterList, nil
}


func (UCS UserCharacterService) Insert(userID string, characterID string) (UserCharacter, error) {
	db := db.GetDB()
	var userCharacter UserCharacter

	uu, err := GenerateUUID()
	if err != nil{
		return userCharacter, err
	}

	userCharacter.UserID = userID
	userCharacter.CharacterID = characterID
	userCharacter.UserCharacterID = uu
	userCharacter.CreatedAt = time.Now()
	userCharacter.ModifiedAt = time.Now()

	if err := db.Create(&userCharacter).Error; err != nil {
		return userCharacter, err
	}

	return userCharacter, nil
}


//Get a UserCharacter
func (UCS UserCharacterService) GetOne(id string) (UserCharacter, error) {
	db := db.GetDB()
	var userCharacter UserCharacter

	if err := db.Where("user_character_id = ?", id).First(&userCharacter).Error; err != nil {
		return userCharacter, err
	}

	return userCharacter, nil
}


//Delete a UserCharacter
func (UCS UserCharacterService) Delete(id string) error {
	db := db.GetDB()
	var userCharacter UserCharacter

	if err := db.Where("user_character_id = ?", id).Delete(&userCharacter).Error; err != nil {
		return err
	}

	return nil
}