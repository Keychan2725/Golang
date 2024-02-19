package main

import (
	"belajargolang/database"
	_ "belajargolang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
