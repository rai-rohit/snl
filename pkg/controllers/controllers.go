package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snakes_ladders/pkg/database"
	"snakes_ladders/pkg/models"
	"snakes_ladders/pkg/utils"
	"strconv"
)

func StartGame(w http.ResponseWriter, r *http.Request) {
	database.DB.Migrator().DropTable(&models.Player{})
	database.DB.Migrator().DropTable(&models.Snake{})
	database.DB.Migrator().DropTable(&models.Ladder{})

	database.DB.AutoMigrate(&models.Player{}, &models.Snake{}, &models.Ladder{})

	err := utils.GetSNl()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error while getting data from DB")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New Game has begun!!"))
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	var p models.Players
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintf(w, "error while decoding - \n %v", err)
		return
	}

	players := make([]models.Player, 0)
	if len(p.Names) != p.Number {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error : Please enter valid number of players.")
		return
	}

	for _, v := range p.Names {
		players = append(players, models.Player{
			Name: v,
		})
	}

	db := database.DB.Table("players").Create(&players)
	if db.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error : Data insertion failed")
		return
	}
	resp, err := json.Marshal(players)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Details of players addeed-\n"))
	w.Write(resp)
}

func DiceRoll(w http.ResponseWriter, r *http.Request) {
	
	id := r.URL.Query().Get("id")
	intId, _ := strconv.Atoi(id)
	player, err := database.GetPlayerById(intId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "error : Enter valid ID")
		return
	}

	players, err := database.GetAllPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error while retrieving data")
		return
	}

	var m = models.Response{
		Message: "Sorry,not your turn",
	}
	resp, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error :", err)
		return
	}

	for _, v := range players {
		if player.NumTurns > v.NumTurns {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(resp)
			return
		}
	}

	x := utils.RandGenerator(6)
	if x != 6 {
		fmt.Fprintf(w, "you got %d on dice roll\n", x)
		player.CurrPos += x
		utils.UpdatePlayerPos(&player)
	} else {
		fmt.Fprintf(w, "you got %d on dice roll\n", x)
		for x%6 == 0 {
			player.CurrPos += x
			utils.UpdatePlayerPos(&player)
			x = utils.RandGenerator(6)
			fmt.Fprintf(w, "you got %d on dice roll\n", x)
		}
		player.CurrPos += x
		utils.UpdatePlayerPos(&player)
	}

	if player.CurrPos > 100 {
		player.CurrPos -= x
	}

	player.NumTurns++
	database.DB.Save(&player).Table("players")

	if player.CurrPos == 100 {
		fmt.Fprintf(w, "%s wins!!\n", player.Name)
	}
}

func GetPlayerPosition(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	intId, err := strconv.ParseInt(id, 0, 0)
	if err !=nil{
		fmt.Fprintf(w,"error :%v",err)
		return
	}

	player, er:= database.GetPlayerById(int(intId))
	if er !=nil{
		fmt.Fprintf(w,"error :%v",er)
		return
	}

	var m = models.Response{
		Message: "Player Postion Details  Is As Follows",
		Data: player,
	}
	resp, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "error :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetAllPlayersPositions(w http.ResponseWriter, r *http.Request) {
	players, err := database.GetAllPlayers()
	if err != nil {
		fmt.Fprintf(w, "error - %v", err)
		return
	}
	fmt.Fprintf(w, "Players Postions Detail Are As Follow:\n%+v", players)
}

func SnlPositions(w http.ResponseWriter, r *http.Request) {
	snakes, ladders, err := database.GetAllSnl()
	if err != nil {
		fmt.Fprintf(w, "error - %v", err)
		return
	}
	fmt.Fprintf(w, " Positions of Snakes :\n %+v \n Positions of Ladders :\n %+v \n", snakes, ladders)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "error while decoding-", err)
	}

	err = database.MatchUserCredentials(user)
	if err != nil {
		fmt.Fprintln(w, "error encountered-", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("username and pass do not match"))
		return
	}

	tokenString, err := utils.GenToken(user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "login successful")

	json.NewEncoder(w).Encode(tokenString)
	if err != nil {
		fmt.Println(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error : %s", err)
		return
	}

	database.DB.AutoMigrate(&models.User{})

	err = database.SaveUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error : %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "registration successful")
}

