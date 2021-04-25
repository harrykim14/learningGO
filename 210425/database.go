package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DBconnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DBconnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DBconnection.Close()
	command := `CREATE TABLE IF NOT EXISTS person(name STRING, age  INT)`
	_, err := DBconnection.Exec(command)
	if err != nil {
		log.Fatalln(err)
	}

	insertCmd := "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DBconnection.Exec(insertCmd, "Jone", 31)
	if err != nil {
		log.Fatalln(err)
	}
	// sqlite> Nancy|20

	updateCmd := "UPDATE person SET age = ? WHERE name = ?"
	_, err = DBconnection.Exec(updateCmd, 25, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
	// sqlite> Nancy|25

	selectCmd := "SELECT * FROM person"
	rows, _ := DBconnection.Query(selectCmd)
	defer rows.Close()
	var pp []Person

	for rows.Next() {
		var p Person
		// err := rows.Scan(&p.Name, &p.Age)
		// if err != nil {
		// 	log.Println(err)
		// }
		pp = append(pp, p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	singleSelectCmd := "SELECT * FROM person where age = ?"
	row := DBconnection.QueryRow(singleSelectCmd, 21)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	deleteCmd := "DELETE FROM person WHERE name = ?"
	_, err = DBconnection.Exec(deleteCmd, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	tableName := "person"
	tableViewCmd := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DBconnection.Query(tableViewCmd)
	defer rows.Close()
	var pp []Person

	for rows.Next() {
		var p Person
		pp = append(pp, p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

}
