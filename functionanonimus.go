package main
import "fmt"

type Blacklist func(string) bool

func registerUser(name string , blacklist Blacklist){
	if blacklist(name){
		fmt.Println("Kamu telah diblokir " , name)
	} else {
		fmt.Println("selamat datang " , name)
	}
} 

func main(){
//  cara 1
	blacklist := func(name string) bool {
		return name == "saru"
	}
	registerUser("chandra" , blacklist)
// cara 2
    registerUser("saru" , func(name string) bool {
		return name == "saru"
	})
 
    
}
