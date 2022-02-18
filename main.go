package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "118.98.64.74"
	port     = 5432
	user     = "postgres"
	password = "S@y@Br0nt0"
	dbname   = "belajar_golang"
)

type User struct {
	ID       string `json:"id" db:"id"`
	Nama     string `json:"nama" db:"username"`
	Password string `json:"password" db:"password"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Query the database, storing results in a []Person (wrapped in []interface{})
	users := []User{}
	err = db.Select(&users, "SELECT * FROM user_ustadho ORDER BY username ASC")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("users: %+v", users)
	fmt.Println("Successfully connected!")
}