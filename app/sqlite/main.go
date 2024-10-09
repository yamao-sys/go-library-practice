package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatalln(err)
	}

	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS persons(name STRING, age INT)`
	// _, error := Db.Exec(cmd)
	// if error != nil {
	// 	log.Fatalln(error)
	// }

	// // データの追加
	// insertCmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, insertErr := Db.Exec(insertCmd, "tarou", 20)
	// if insertErr != nil {
	// 	log.Fatalln(insertErr)
	// }

	// // データの更新
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// _, updateErr := Db.Exec(cmd, 25, "tarou")
	// if updateErr != nil {
	// 	log.Fatalln(updateErr)
	// }

	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, insertErr := Db.Exec(cmd, "hanako", 19)
	// if insertErr != nil {
	// 	log.Fatalln(insertErr)
	// }

	// // 特定のデータを取得
	// cmd := "SELECT * FROM persons WHERE age = ?"
	// // QueryRow 1レコードの取得
	// row := Db.QueryRow(cmd, 25)
	// var p Person
	// scanErr := row.Scan(&p.Name, &p.Age)
	// if scanErr != nil {
	// 	if scanErr == sql.ErrNoRows {
	// 		log.Println("No Rows")
	// 	} else {
	// 		log.Fatalln(scanErr)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// // 複数データの取得
	// cmd := "SELECT * FROM persons"
	// // Queryは条件に合うものを全て取得
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	scanErr := rows.Scan(&p.Name, &p.Age)
	// 	if scanErr != nil {
	// 		log.Println(scanErr)
	// 	}
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, deleteErr := Db.Exec(cmd, "hanako")
	if deleteErr != nil {
		log.Fatalln(deleteErr)
	}
}
