package main

import "fmt"

type ValidationError struct {
	message string
}

func (v *ValidationError) Error() string {
	return v.message
}

type NotFOundError struct {
	message string
}

func (n *NotFOundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"Id Anda Tidak Ditemukan"}
	}
	if id != "chandra" {
		return &NotFOundError{"Nama Anda Tidak Ditemukan"}
	}

	return nil
}
func main() {

	err := SaveData("yan", nil)

	if err != nil {
		//if ValidationError, ok := err.(*ValidationError); ok {
		//	fmt.Println("Validasi erro : ", ValidationError.Error())
		//} else if NotFOundError, ok := err.(*NotFOundError); ok {
		//	fmt.Println("User tidak ditemukan :", NotFOundError.Error())
		//} else {
		//	fmt.Println("Error Gj", err.Error())
		//}
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("Validasi error : ", finalError.Error())
		case *NotFOundError:
			fmt.Println("Notfound error : ", finalError.Error())
		default:
			fmt.Println("Error Gj", finalError.Error())
		}

	} else {
		fmt.Println("Berhasil")
	}
}
