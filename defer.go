package main
import "fmt"
// untuk membaca func sesudah func berjalan

func logging(){
	fmt.Println("Berhasil Login")
}

func runAplication(){
	defer logging()

	fmt.Println("Run Aplicaton")
}

func main(){

	runAplication()
}