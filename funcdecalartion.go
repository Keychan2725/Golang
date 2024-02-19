package main
import "fmt"

type Filter func(string) string
type number func(int) int

func sayHello(name string , filter Filter ) {
	filteredName := filter(name)
	fmt.Println("Hello " , filteredName)

 

}

func spamFilter(name string ) string{
	if name == "anjing"{
		return "xixixi"
	} else {
		return name
	}

}
 
func main(){
	sayHello("eko" , spamFilter)

	filter := spamFilter
	sayHello("anjing" ,filter)
 
}