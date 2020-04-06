package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

func CityList() {
	// sql.Open("mysql", "user:password@tcp(container-name:port)/dbname")
	db, err := sql.Open("mysql", "dbuser:dbpassword@tcp(go_db:3306)/world")
	log.Println("Connected to mysql.")

	//接続でエラーが発生した場合の処理
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM city")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var city City
		err := rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(city.ID, city.Name, city.CountryCode, city.District, city.Population)

	}
}
