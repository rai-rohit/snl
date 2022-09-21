package database

import (
	"fmt"
	"snakes_ladders/pkg/config"
	"snakes_ladders/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *config.Config) error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.Name)
	d, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = d
	return nil
}

func GetPlayerById(id int) (models.Player, error) {
	var player models.Player
	db := DB.Table("players").First(&player, id)
	if db.Error != nil {
		return player, db.Error
	}
	return player, nil
}
func GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	db := DB.Find(&players)
	if db.Error != nil {
		return players, db.Error
	}
	return players, nil
}

func GetAllSnl() ([]models.Snake, []models.Ladder, error) {
	var snakes []models.Snake
	var ladders []models.Ladder
	db := DB.Find(&snakes)
	if db.Error != nil {
		return snakes, nil, db.Error
	}
	db = DB.Find(&ladders)
	if db.Error != nil {
		return nil, ladders, db.Error
	}
	return snakes, ladders, nil
}

func SaveUser(u models.User) error {
	db := DB.Table("users").Create(&u)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func MatchUserCredentials(u models.User) error {
	var password string
	db := DB.Table("users").Select("password").Where("email = ?", u.Email).Scan(&password)
	if db.Error != nil {
		return db.Error
	}
	if u.Password != password {
		return fmt.Errorf("password match failed %s", u.Password)
	}
	return nil
}
