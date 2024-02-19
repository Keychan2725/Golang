package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, er := pembagian(330, 0)
	if er == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Erorrrr", er.Error())
	}

}
