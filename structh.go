package main 
import "fmt"

type Customer struct{
	Name          string
    Alamat        string   
	Umur          int
	
}
func (customer Customer) SayHello() {

	fmt.Println("Haloo , Nama sayaa " , customer.Name , ". Umur saya " , customer.Umur )
}
func main(){

	Chan := Customer{Name: "Wanderer"}
	Chan.SayHello()


	var chandra Customer
	chandra.Name = "Wanderer Chandra"
	chandra.Alamat = "Karangreja"
	chandra.Umur = 15

	fmt.Println(chandra)
}