package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 構造体からJSONテキストへの変換

type A struct{}

// JSONの表記を`json:"xx"`で指定している
type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
	B       int       `json:"-"`           // 常にomit
	C       string    `json:"C,omitempty"` // 空値であればomit
	D       *A        `json:"D,omitempty"` // 構造体のomitemptyはポインタに対して有効
}

// カスタムMarshal
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

// カスタムUnmarshal
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "exaple@example.com"
	u.Created = time.Now()
	u.B = 10
	u.C = ""

	// Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	// JSONから構造体への変換

	u2 := new(User)

	// Unmarshal JSONを構造体データに変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
