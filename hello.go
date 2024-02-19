package main 

import "fmt"

// func SayChandra(firstname string , lastname string){
// 	fmt.Println("helloo" , firstname , lastname)
// }

// func getHello( name string) string {
	
// 	Hello := "Hello " + name
// 	return Hello
// }
// func main(){
 
//     result := getHello("chandra")
// 	fmt.Println(result)
// 	// SayChandra("Wanderer", "Chandra")
// }

// func MultipleFUnction() (string , string) {
//   return "Chandra" , "Wanderer"
// }

func variadic(numbers ...int) int {

	total := 0
	
	for _, number := range numbers{

		total += number
	}
	return total
}

func main(){

	fmt.Println(variadic(12,12,12,12,12))
	
	numbers := []int{20,20,20}
	fmt.Println(variadic(numbers...))
	// namadepan , namabelakang := MultipleFUnction()
	// fmt.Println(namadepan , namabelakang)
}