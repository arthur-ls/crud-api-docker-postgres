package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arthur-ls/project1/database"
	"github.com/arthur-ls/project1/entities"
)

var comp entities.Competition
var db, err = database.InitDB()

func CreateTable() {
	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS competition(LeagueID INT NOT NULL,CountryID INT NOT NULL,Name VARCHAR(100) NOT NULL);"); err != nil {
		panic(err)
	}
}

func GetAll() ([]entities.Data, error) {
	defer db.Close()

	var competition []entities.Data

	sqlStatement := `SELECT * FROM competition`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var competition_var entities.Data

		err = rows.Scan(&competition_var.LeagueID, &competition_var.CountryID, &competition_var.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		competition = append(competition, competition_var)
	}

	return competition, err
}

func GetCompetitionById(id int64) (entities.Data, error) {
	defer db.Close()
	var i entities.Data
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}
	sqlGetAll := `SELECT * FROM competition WHERE LeagueID = $1`
	row := db.QueryRow(sqlGetAll, id)
	err = row.Scan(&i.LeagueID, &i.CountryID, &i.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return i, nil
	case nil:
		return i, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return i, err
}

func DeleteCompetitionById(id int64) int64 {
	defer db.Close()

	sqlStatement := `DELETE FROM competition WHERE LeagueID=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func InsertCompetitionById(competition entities.Data) int64 {
	defer db.Close()

	sqlStatement := `INSERT INTO competition (LeagueID, CountryID, Name) VALUES ($1, $2, $3) RETURNING LeagueID`

	var id int64

	err = db.QueryRow(sqlStatement, competition.LeagueID, competition.CountryID, competition.Name).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func UpdateCompetitionById(id int64, comp entities.Data) int64 {
	defer db.Close()

	sqlStatement := `UPDATE competition SET CountryID=$2, Name=$3 WHERE LeagueID=$1`

	res, err := db.Exec(sqlStatement, id, comp.CountryID, comp.Name)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
