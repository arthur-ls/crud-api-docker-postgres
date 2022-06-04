package database

func DatabaseParams() (string, int, string, string, string) {
	host := "localhost"
	port := 5432
	user := "docker"
	password := "docker"
	dbname := "rentx"

	return host, port, user, password, dbname
}
