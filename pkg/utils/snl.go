package utils

import (
	"fmt"
	"math/rand"
	"snakes_ladders/pkg/database"
	"snakes_ladders/pkg/models"
	"time"
)

func GetSnakes() error {
	xs := make([]models.Snake, 5)

	var i uint
	for i < 5 {
		xs[i].Head = RandGenerator(99)
		xs[i].Tail = RandGenerator(99)
		if Validator(xs[i].Head, xs[i].Tail) {
			xs[i].Name = fmt.Sprintf("snake %d", i+1)
			i++
		}
	}

	if !uniqueSnakes(xs) {
		GetSnakes()
	}

	db := database.DB.Table("snakes").Create(&xs)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func GetLadders() error {
	xl := make([]models.Ladder, 5)

	var i uint
	for i < 5 {
		xl[i].High = RandGenerator(99)
		xl[i].Low = RandGenerator(99)
		if Validator(xl[i].High, xl[i].Low) {
			xl[i].Name = fmt.Sprintf("ladder %d", i+1)
			i++
		}
	}
	if !uniqueLadders(xl) {
		GetLadders()
	}
	db := database.DB.Table("ladders").Create(&xl)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func uniqueSnakes(snakes []models.Snake) bool {
	unique := make(map[models.Snake]bool)

	for _, v := range snakes {
		unique[v] = true
	}
	return len(unique) == len(snakes)
}

func uniqueLadders(ladders []models.Ladder) bool {
	unique := make(map[models.Ladder]bool)

	for _, v := range ladders {
		unique[v] = true
	}
	return len(unique) == len(ladders)
}

func RandGenerator(x int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)
	e := (rand.Intn(x) + 1)
	return e
}

func Validator(high, low int) bool {
	if high < low {
		return false
	} else if high/10 == low/10 {
		return false
	}
	return true
}

func UpdatePlayerPos(player *models.Player) {
	snakes, ladders, err := database.GetAllSnl()
	if err != nil {
		fmt.Println("error - %w", err)
		return
	}

	for _, snake := range snakes {
		if player.CurrPos == snake.Head {
			player.CurrPos = snake.Tail
		}
	}

	for _, ladder := range ladders {
		if player.CurrPos == ladder.Low {
			player.CurrPos = ladder.High
		}
	}
}
