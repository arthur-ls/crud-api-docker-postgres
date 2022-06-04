package api

import (
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func ApiData() ([]byte, error) {
	url := "https://app.sportdataapi.com/api/v1/soccer/leagues"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("apikey", "fb0e7c40-db81-11ec-ac9b-6fd67f3c9f94")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	requestBody, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return requestBody, nil
}
