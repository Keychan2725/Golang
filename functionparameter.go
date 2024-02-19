package main
import "fmt"

func sayHelloWithFilter(name string , filter func(string) string){
	fmt.Println("Hello " + filter(name) )
}

func spamFilter(name string) string{
	if name == "Babi" {
		return "..."
	} else {
		return name 
	}

}
func main(){
	sayHelloWithFilter("eko" , spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Babi" ,filter)
}
