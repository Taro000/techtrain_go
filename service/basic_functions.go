package service

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func GenerateUUID() (string, error) {
	var uu string
	u, err := uuid.NewRandom()
	if err != nil {
		return uu , err
	}
	uu = u.String()
	return uu, nil
}


func GenerateLotteryList(characters []Character) []Character {
	var lotteryList []Character

	for _, v := range characters{
		for i := 0; i < v.Probability; i++{
			lotteryList = append(lotteryList, v)
		}
	}
	return lotteryList
}


func Lottery(lotteryList []Character) Character {
	rand.Seed(time.Now().UnixNano())
	return lotteryList[rand.Intn(len(lotteryList))]
}