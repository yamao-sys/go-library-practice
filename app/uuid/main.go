package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println(" ", uuidObj2.String())

	fmt.Println("version3 NewMD5 --")
	for i := 0; i < 10; i++ {
		uuidObj, _ := uuid.NewUUID()
		data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
		uuidObj2 := uuid.NewMD5(uuidObj, data)
		fmt.Println("  ", uuidObj2.String())
	}

	fmt.Println("version5 NewSHA1 --")
	for i := 0; i < 10; i++ {
		uuidObj, _ := uuid.NewUUID()
		data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
		uuidObj2 := uuid.NewSHA1(uuidObj, data)
		fmt.Println("  ", uuidObj2.String())
	}
}
