package models

type Players struct {
	Names  []string `json:"names"`
	Number int      `json:"number"`
}

type Snake struct {
	Name string `json:"name"`
	Head int    `json:"head"`
	Tail int    `json:"tail"`
}

type Ladder struct {
	Name string `json:"name"`
	High int    `json:"high"`
	Low  int    `json:"low"`
}

type Database struct {
	Dbtype   string `yaml:"dbtype"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
}

type Player struct {
	Id       int    `json:"id" gorm:"unique,autoIncrement"`
	Name     string `json:"name"`
	CurrPos  int    `json:"currpos"`
	NumTurns int    `json:"numturns"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"string,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type Simulation struct {
// 	Id          int    `json:"id"`
// 	NumRoll     int    `json:"numroll"`
// 	Climb       int    `json:"climb"`
// 	Slide       int    `json:"slide"`
// 	Turn        string `json:"turn"`
// 	LuckyRoll   int    `json:"luckyroll"`
// 	UnluckyRoll int    `json:"unluckyroll"`
// 	IsWinner    bool   `json:"iswinner"`
// }
