package main 
import "fmt"

type Man struct{
	Name string
}

func (man *Man)  Married(){
	man.Name = "Mr. " +  man.Name
}

func main(){

	chandra := Man{"Chandra"}

	chandra.Married()

	fmt.Println(chandra.Name)
}