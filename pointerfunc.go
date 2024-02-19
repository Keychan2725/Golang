package main
import "fmt"

type Alamat struct{
	Kota, Provinsi, Negara string
}
//  opertaor * digunakan untuk mengakses kedalama struct
func GantiNegara(alamat *Alamat){
    alamat.Negara ="Kudus"
}

func main(){

	alamat := Alamat{"Pekalongan" ,"Kalimantan" , ""}
	// untuk membuat pointer ke strutc yang sama
	GantiNegara(&alamat)

	fmt.Println(alamat)
}