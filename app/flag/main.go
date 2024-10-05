package main

import (
	"flag"
	"fmt"
)

func main() {
	// コマンドラインのオプション処理
	// コマンドラインを処理するサンプル
	// go run main.go -n 20 -m message -x

	// オプションの値を格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	// IntVar 整数のオプション
	flag.IntVar(&max, "n", 32, "処理数の最大値") // オプション名 n, デフォルト値 32 オプション説明文 処理数の最大値
	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理メッセージ") // オプション名 m, デフォルト値 '' オプション説明文 処理メッセージ
	// BoolVar bool型のオプション コマンドラインに与えられたらtrue なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション") // オプション名 x, デフォルト値 false オプション説明文 拡張オプション

	// コマンドラインのParse
	flag.Parse() // ↑のflagで設定したオプションに値が入る

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション =", x)
}
