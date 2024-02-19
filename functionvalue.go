package main

import "fmt"

func sayGoodBye(name string)string{
	return "Sayonaraa "+ name
}

func main(){
	goodbye := sayGoodBye
	fmt.Println(goodbye("chandra"))
}
