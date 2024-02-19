package main 
import "fmt"

func recursivefactorial(value int) int {

	if value == 1{
		return 1
	} else {
		return value * recursivefactorial(value-1)
	}
}

func main(){

	fmt.Println(recursivefactorial(10))
}