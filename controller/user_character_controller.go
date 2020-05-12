package controller

import (
	_ "bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "io/ioutil"
	_ "strconv"
	_ "techtrain_go_api/models"
	"techtrain_go_api/service"
)


//Controller of User
type UserCharacterController struct {}

type gottenCharacter map[string]string


func (UCC UserCharacterController) Gacha(context *gin.Context) {
	token := context.Request.Header.Get("x-token")
	var cs service.CharacterService
	var us service.UserService
	var ucs service.UserCharacterService
	var result []gottenCharacter
	var gacha_json map[string]int

	user, err := us.GetOne(token)
	if err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	}

	characters, err := cs.GetAll()
	if err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	}

	//buf := make([]byte, 2048)
	//n, _ := context.Request.Body.Read(buf)
	//b := string(buf[0:n])
	//context.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))
	//fmt.Println(b)

	buf := make([]byte, 2048)
	n, _ := context.Request.Body.Read(buf)
	b := buf[0:n]

	if err := json.Unmarshal(b, &gacha_json); err != nil{
		context.AbortWithStatus(400)
		fmt.Println(err)
	}

	//bodyCopy := new(bytes.Buffer)
	//_, err = io.Copy(bodyCopy, context.Request.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//bodyData := bodyCopy.Bytes()
	//context.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))

	//if err := context.BindJSON(json); err != nil{
	//	context.AbortWithStatus(400)
	//	fmt.Println(err)
	//}

	if !(gacha_json["times"] >= 1) {
		context.AbortWithStatus(400)
		fmt.Println("times need to be over 1(int).")
	}

	//times, _ := strconv.Atoi(json.Times)
	times := gacha_json["times"]

	for i := 0; i < times; i++ {
		lotteryList := service.GenerateLotteryList(characters)
		character := service.Lottery(lotteryList)

		_, err = ucs.Insert(user.ID, character.ID)
		result = append(result, gottenCharacter{"characterID": character.ID, "name": character.Name})

		if err != nil {
			context.AbortWithStatus(400)
			fmt.Println(err)
		}
	}
	context.JSON(201, result)
}