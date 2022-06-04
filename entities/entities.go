package entities

type Competition struct {
	Query []interface{} `json:"query"`
	Data  []struct {
		LeagueID  int    `json:"league_id"`
		CountryID int    `json:"country_id"`
		Name      string `json:"name"`
	} `json:"data"`
}

type Data struct {
	LeagueID  int    `json:"league_id"`
	CountryID int    `json:"country_id"`
	Name      string `json:"name"`
}
