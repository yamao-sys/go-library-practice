package main

import (
	"fmt"
	"sync"
	"time"
)

// Goの非同期処理をサポートする同期処理のための機能が備わるsync

// ミューテックスによる同期処理

// レースコンディション
// 並列動作する複数のプロセスが同一のリソースに同時アクセスした時に予想外の処理結果が生じてしまうこと

var st struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
// mutexによる排他制御
var mutex *sync.Mutex

// 同期すべきブロックの検討が不十分でmutexを入れまくると、せっかくの非同期の利便性を損ねてしまう
// → どのような範囲と単位で同期化するのかを検討したいところ

// ゴルーチンの終了を待ち受ける

func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)

	// アンロック
	mutex.Unlock()
}

func main() {
	// mutex = new(sync.Mutex)

	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		for i := 0; i < 1000; i++ {
	// 			UpdateAndPrint(i)
	// 		}
	// 	}()
	// }

	// for {
	// }

	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	// 待ち受けるゴルーチンの数は3
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() // 完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() // 完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() // 完了
	}()

	// ゴルーチンの完了を待ち受ける
	// Doneが3つ完了するまで待つ
	wg.Wait()
}
