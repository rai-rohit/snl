package utils

import (
	"fmt"
	"math/rand"
	"snakes_ladders/pkg/database"
	"snakes_ladders/pkg/models"
	"time"
)

func GetSNl() error {

	xs := []models.Snake{}
	xl := []models.Ladder{}

	for {
		xl = getLadders()
		xs = getSnakes()

		// cross checking 
		if uniqueSnL(xs, xl) {
			break
		}
	}

	db := database.DB.Table("snakes").Create(&xs)
	if db.Error != nil {
		return db.Error
	}

	db = database.DB.Table("ladders").Create(&xl)
	if db.Error != nil {
		return db.Error
	}
	return nil

}

func getSnakes() []models.Snake {
	xs := make([]models.Snake, 5)

	var i uint
	for i < 5 {
		xs[i].Head = RandGenerator(99)
		xs[i].Tail = RandGenerator(99)
		if Validator(xs[i].Head, xs[i].Tail) {
			xs[i].Name = fmt.Sprintf("snake-%d", i+1)
			i++
		}
	}

	if !uniqueSnakes(xs) {
		xs = getSnakes()
	}

	return xs
}

func getLadders() []models.Ladder {
	xl := make([]models.Ladder, 5)

	var i uint
	for i < 5 {
		xl[i].High = RandGenerator(99)
		xl[i].Low = RandGenerator(99)
		if Validator(xl[i].High, xl[i].Low) {
			xl[i].Name = fmt.Sprintf("ladder-%d", i+1)
			i++
		}
	}
	if !uniqueLadders(xl) {
		xl = getLadders()
	}
	return xl
}

// no two snakes should have same head
func uniqueSnakes(snakes []models.Snake) bool { 
	unique := make(map[int]bool)
	for _, v := range snakes {
		unique[v.Head] = true
	}
	return len(unique) == len(snakes)
}
//no two ladders should have same start point
func uniqueLadders(ladders []models.Ladder) bool {
	unique := make(map[int]bool)
	for _, v := range ladders {
		unique[v.Low] = true
	}
	return len(unique) == len(ladders)
}

func uniqueSnL(snakes []models.Snake, ladders []models.Ladder) bool {
	// any snake's head should be at the bottom point of any ladder
	unique := make(map[int]bool)
	for i, v := range ladders {
		unique[v.Low] = true
		unique[snakes[i].Head] = true
	}
	if len(unique) != len(snakes)+len(ladders) {
		return false
	}
	// any snake's head should not be at the top point of any ladder
	unique2 := make(map[int]bool)
	for i, v := range ladders {
		unique2[v.High] = true
		unique2[snakes[i].Head] = true
	}
	if len(unique2) != len(snakes)+len(ladders) {
		return false
	}
	// any snake's tail should not be at the bottom of any ladder
	unique3 := make(map[int]bool)
	for i, v := range ladders {
		unique3[v.Low] = true
		unique3[snakes[i].Tail] = true
	}
	return len(unique3) == len(snakes)+len(ladders)

}

func RandGenerator(x int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)
	return (rand.Intn(x) + 1)
}

func Validator(top, bottom int) bool {
	if top < bottom {
		return false
	} else if top/10 == bottom/10 {
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
