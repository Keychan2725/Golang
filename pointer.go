package main
import "fmt"

type Alamat struct{
	Kota, Provinsi, Negara string
}

func main(){
	alamat1 := Alamat{"Purbaligga" ,"Jawa Tengah" ,"Indonesia"}
	alamat2 := alamat1 //gnti alamat1 dengan &alamat1 untuk mengganti struct alamat1

	alamat2.Negara = "Jepang"

	fmt.Println(alamat1) //tidak beruba
	fmt.Println(alamat2) //berubah

}