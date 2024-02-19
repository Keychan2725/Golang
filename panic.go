package main
import "fmt"
// mendeteksi erorr
func endApp(){
	fmt.Println("Aplikasi Tidak berjalan")
	message := recover()

	fmt.Println("Terjadi eror ," , message)
}

func Panic( eror bool){
	defer endApp()
	if eror {
		panic("Kamuu wibu")
	}

}

func main(){
	Panic(true)
}
