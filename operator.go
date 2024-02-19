package main 
import "fmt"

type Alamat struct{
	Kota, Provinsi, Negara string
}

func main(){
	var alamat1 := Alamat{"Purbaligga" ,"Jawa Tengah" ,"Indonesia"}
	var alamat2  := &alamat1 //gnti alamat1 dengan &alamat1 untuk mengganti struct alamat1

	alamat2.Negara = "Jepang"
	alamat2 = &Alamat{"Jakarta" , "Jabar" , "Indonesia"}

	//operator asteris
	// bisa juga
	// alamat1 := new(Alamat)
	// 	alamat1 := Alamat{"Purbaligga" ,"Jawa Tengah" ,"Indonesia"}
	// 	alamat2  := &alamat1 //gnti alamat1 dengan &alamat1 untuk mengganti struct alamat1

	//    alamat2.Negara = "Jepang"
	//    *alamat2 = Alamat{"Jakarta" , "Jabar" , "Indonesia"}


	fmt.Println(alamat1) //tidak beruba
	fmt.Println(alamat2) //berubah
}