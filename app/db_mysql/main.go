package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_library_practice"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	Db, connectErr := sql.Open(DBMS, CONNECT)
	if connectErr != nil {
		log.Fatalln(connectErr)
	}
	defer Db.Close()

	// // C
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // R
	// cmd := "SELECT * FROM persons WHERE age = ?"
	// // QueryRow 1レコード取得
	// row := Db.QueryRow(cmd, 20)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	// データがなかったら
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No Rows")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "SELECT * FROM persons"
	// // Queryは条件に合うものを全て取得
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()

	// // structを作成
	// var pp []Person
	// // 取得したデータをループでsliceに追加
	// for rows.Next() {
	// 	var p Person
	// 	// scanデータ追加
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	// 1つずつエラーハンドリング
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// // まとめてエラーハンドリング
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // 表示
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// // U
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// _, err := Db.Exec(cmd, 25, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// D
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
