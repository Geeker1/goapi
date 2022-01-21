package db_test

import (
	"log"
	"testing"

	"github.com/Geeker1/goapi/db"
)

// Before running test, export DB_PATH (full path to local json file for database.)


func TestDB(t *testing.T)  {
	apiDB, err := db.InitDB()

	if err != nil {
		log.Panic("Error initializing local database", err)
	}

	t.Run("test all club data is fetched", func(t *testing.T) {
		clubs := apiDB.GetClubs()

		if len(clubs) != 12 {
			t.Errorf("Club list length different from expected.")
		}

		club := clubs[0]

		if club.Name != "Chelsea FC" {
			t.Errorf("Club data not found.")
		}
	})

	t.Run("test single club data is retrieved", func(t *testing.T) {
		code := "RMD"

		club, status := apiDB.GetClub(code)

		if status == false {
			t.Errorf("Error fetching club data")
		}

		if club.Name != "Real Madrid" {
			t.Errorf("Club code does not match data fetched")
		}
	})

	t.Run("test single club function throws error on wrong code", func(t *testing.T) {
		code := "JISDJS"

		_, status := apiDB.GetClub(code)

		if status == true {
			t.Errorf("Expected error, got positive status")
		}

	})
}

