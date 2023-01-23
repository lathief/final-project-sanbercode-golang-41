package main

import (
	"bioskop/database"
	"bioskop/routers"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	var PORT string
	var args = flag.String("env", "dev", "type your APP_ENV")
	if strings.Compare(*args, "prod") == 1 {
		err := godotenv.Load("config/app.env")
		if err != nil {
			fmt.Println("Failed Load file env")
			panic(err)
		} else {
			fmt.Println("Successfully Load file env")
		}
		PORT = ":" + os.Getenv("PORT")

	} else {
		err := godotenv.Load("config/.env")
		if err != nil {
			fmt.Println("Failed Load file env")
			panic(err)
		} else {
			fmt.Println("Successfully Load file env")
		}
		PORT = ":8081"
	}
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, db_name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed connected to database")
		panic(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	database.DbMigrate(db)

	defer db.Close()

	routers.StartServer().Run(PORT)
}
