package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:13306)/goweb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		query := `
			create table if not exists users (
				id int auto_increment,
				username text not null,
				password text not null,
				created_at datetime,
				primary key (id)
			);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		// insert one user
		var username = "lihua"
		var password = "lihua666"
		var created_at = time.Now()

		result, err := db.Exec(`insert into users(username, password, created_at) values(?, ?, ?)`, username, password, created_at)
		if err != nil {
			log.Fatal(err)
		}
		userId, _ := result.LastInsertId()
		fmt.Println(userId)
	}

	{
		// query a single user
		query := `
		select id, username, password from users where id = ?`

		var (
			id       int
			username string
			password string
		)
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, password)
	}

	{
		// query all users

		type user struct {
			id       int
			username string
			password string
		}

		rows, _ := db.Query(`select id, username, password from users`)
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		fmt.Printf("%#v", users)

	}

	{
		// delete a user
		if _, err := db.Exec(`delete from users where id = ?`, 1); err != nil {
			log.Fatal(err)
		}

	}

}
