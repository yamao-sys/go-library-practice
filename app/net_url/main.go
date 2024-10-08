package main

import (
	"fmt"
	"net/url"
)

func main() {
	// // URLを解析
	// u, _ := url.Parse("http://example.com/search?a=1&b=2#top") // URL型の構造体を返す
	// fmt.Println(u.Scheme)
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)
	// fmt.Println(u.RawQuery)
	// fmt.Println(u.Fragment)

	// fmt.Println(u.Query())

	// URLを生成
	url := &url.URL{} // URL型の構造体のポインタを生成
	url.Scheme = "https:"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Golang")

	url.RawQuery = q.Encode()

	fmt.Println(url)
}
