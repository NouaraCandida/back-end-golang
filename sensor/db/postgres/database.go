package database

import(
	"database/sql"
	"fmt"
)

type Config struct{
	Host		string
	Port 		string
	User		string
	Password 	string
	Database	string
}

func getConnect(config *Config) string {
	return fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=disable")
}

func main() {
	//back
	config := &Config{
		Host: 
		Port:
		User: postgres
		Password: postgres
		Database: sensor
	}
	db, err := sql.Open("nrpostgres", getConnect())

}