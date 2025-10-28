package main

import "fmt"

func sum_s(words ...string) {
	for _, word := range words {
		fmt.Printf("%s ", word);
	}
}

func main(){
	sum_s("Hello", "AI", "and", "Bashnya!");
}
