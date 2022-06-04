package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthur-ls/project1/api"
	"github.com/arthur-ls/project1/database"
	"github.com/arthur-ls/project1/entities"
	route "github.com/arthur-ls/project1/routes"
	"github.com/arthur-ls/project1/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var comp entities.Competition

func InsertAll() {
	db, err := database.InitDB()
	sqlInsert := `INSERT INTO public.competition(LeagueID, CountryID, Name) VALUES ($1, $2, $3)`
	for i := 0; i < len(comp.Data); i++ {
		fmt.Printf("\nInserting Row %v...", i)
		_, err = db.Exec(sqlInsert, comp.Data[i].LeagueID, comp.Data[i].CountryID, comp.Data[i].Name)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("\nRow %v Inserted Successfully!", i)
		}
	}
}

func main() {
	database.InitDB()
	services.CreateTable()
	reqBody, _ := api.ApiData()
	json.Unmarshal(reqBody, &comp)
	InsertAll()
	router := mux.NewRouter()
	route.Route(router)

	log.Println(fmt.Sprintf("Starting Server on port 3000"))
	http.ListenAndServe(":3000", router)
}
