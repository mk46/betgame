package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var cl User

// func main() {
// 	var choice int
// 	db := connectPostgresDB()
// 	for {
// 		fmt.Println("Choose\n1.Insert data\n2.Read data\n3.Update data\n4.Delete data\n5.Exit")
// 		fmt.Scan(&choice)
// 		switch choice {
// 		case 1:
// 			Insert(db)
// 		case 2:
// 			Read(db)
// 		case 3:
// 			Update(db)
// 		case 4:
// 			Delete(db)
// 		case 5:
// 			os.Exit(1)
// 		}
// 	}
// }

// CONNECT DB

//before connecting you have to create a database and a table in psql shell (just a base code improve these code as well as you need)

func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=postgres password='mysecretpassword' host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// INSERT

func Insert(db *sql.DB, user User) {
	_, err := db.Exec("INSERT INTO  users(name,email,phone,address) VALUES($1,$2,$3,$4)", user.Name, user.Email, user.Phone, user.Address)

	if err != nil {
		log.Println("failed to insert user", err)
	} else {
		log.Println("user inserted")
	}

}

// func insertIntoPostgres(db *sql.DB, id int, name, domain string) {
// 	_, err := db.Exec("INSERT INTO  user(name,domain) VALUES($1,$2,$3)", id, name, domain)

// 	_, err := db.Exec(table)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("value inserted")
// 	}
// }

// READ

func Read(db *sql.DB, user *User) {
	rows, err := db.Query("SELECT * FROM users where phone=$1", user.Phone)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Name  Email  Phone  Address")
		for rows.Next() {
			rows.Scan(&cl.Name, &cl.Email, &cl.Phone, &cl.Address)
			log.Printf("%s - %s - %s  - %s\n", cl.Name, cl.Email, cl.Phone, cl.Address)
		}

	}
}

// UPDATE

func Update(db *sql.DB, user User) {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2, address=$3  WHERE phone=$4", user.Name, user.Email, user.Address, user.Phone)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data updated")
	}
}

// DELETE

func Delete(db *sql.DB, user User) {
	_, err := db.Exec("DELETE FROM students WHERE phone=$1", user.Phone)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data deleted")
	}
}

// completed ?? now do it in GROM
