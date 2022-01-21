package db

import (
	"encoding/json"
	"log"
	"os"
)

var (
	DB_PATH = os.Getenv("DB_PATH")
	API_DB  = &DB{
		data: make(map[string]Club),
	}
)

type Club struct {
	Name     string `json:"name"`
	Goals    int    `json:"goals"`
	Avatar   string `json:"avatar"`
	Trophies int    `json:"trophies"`
	League   string `json:"league"`
	Position int    `json:"position"`
	Code     string `json:"code"`
}

type DB struct {
	data map[string]Club
}

func InitDB() (*DB, error) {
	err := API_DB.LoadDB()
	return API_DB, err
}

func GetDB() *DB {
	return API_DB
}

func (db *DB) LoadDB() error {
	dat, err := os.ReadFile(DB_PATH)
	if err != nil {
		log.Println("Error opening DB file")
		return err
	}

	var clubs []Club
	jerr := json.Unmarshal(dat, &clubs)
	if jerr != nil {
		log.Println("Unable to parse data")
		return jerr
	}

	for _, club := range clubs {
		db.data[club.Code] = club
	}

	return nil
}

func (db *DB) GetClub(code string) (Club, bool) {
	club, err := db.data[code]
	return club, err
}

func (db *DB) GetClubs() []Club {
	values := make([]Club, 0, len(db.data))
	for _, club := range db.data {
		values = append(values, club)
	}
	return values
}
