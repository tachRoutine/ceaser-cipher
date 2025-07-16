package main

import (
	"flag"
	"fmt"
)

func main(){
	numberOfSubs := flag.Int("sub",1,"number of substitutions of the Caesar cipher")
	content:= flag.String("content","Hello World!","content to be encrypted")
	fmt.Println("Number of substitutions:", *numberOfSubs)
	fmt.Println("Content to be encrypted:", *content)
	flag.Parse()
}