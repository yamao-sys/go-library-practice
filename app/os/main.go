package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("Start")

	// _, err := os.Open("b.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length=%d\n", len(os.Args))
	// // os.Argsの中身を全て表示
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// ファイル操作
	// f, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// ファイル作成
	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang\n"), 6)
	// f.Seek(0, io.SeekEnd)
	// f.WriteString("Yeah")

	// ファイル読み込み
	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// bs := make([]byte, 128)

	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// bs2 := make([]byte, 128)
	// nm, err := f.ReadAt(bs2, 10)
	// fmt.Println(nm)
	// fmt.Println(string(bs2))

	// OpenFile
	// O_RDONLY 読み取り専用
	// O_WRONLY 書き込み専用
	// O_RDWR 読み書き可能
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_RDONLY, 0666) // 第二引数は複数|で指定できる
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
