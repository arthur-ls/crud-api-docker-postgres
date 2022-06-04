package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/arthur-ls/project1/entities"
	"github.com/arthur-ls/project1/services"
	"github.com/gorilla/mux"
)

var comp entities.Competition

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	competition, err := services.GetAll()

	if err != nil {
		log.Fatalf("Unable to get all competitions. %v", err)
	}

	json.NewEncoder(w).Encode(competition)
}

func HandleGetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	competition, err := services.GetCompetitionById(int64(id))

	if err != nil {
		log.Fatalf("Unable to get competition. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(competition)
}

func HandleDeleteID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := services.DeleteCompetitionById(int64(id))

	msg := fmt.Sprintf("Competition deleted successfully. Total rows/record affected %v", deletedRows)

	json.NewEncoder(w).Encode(msg)
}

func HandleInsertByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	var competition entities.Data

	err := json.NewDecoder(r.Body).Decode(&competition)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := services.InsertCompetitionById(competition)

	res := entities.Data{
		LeagueID:  int(insertID),
		CountryID: competition.CountryID,
		Name:      competition.Name,
	}

	msg := fmt.Sprintf("Row Inserted Successfully")

	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(msg)
}

func HandleUpdateID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to conver the string into int. %v", err)
	}

	var competition entities.Data

	err = json.NewDecoder(r.Body).Decode(&competition)

	if err != nil {
		log.Fatalf("Unable to decode the request Body. %v", err)
	}

	updatedRows := services.UpdateCompetitionById(int64(id), competition)

	msg := fmt.Sprintf("Competition updated successfully. Total rows/record affected %v", updatedRows)

	res := entities.Data{
		LeagueID:  id,
		CountryID: competition.CountryID,
		Name:      competition.Name,
	}

	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(msg)
}
