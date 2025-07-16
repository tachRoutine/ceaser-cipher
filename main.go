package main

import "flag"

func main(){
	numberOfSubs := flag.Int("sub",1,"number of substitutions of the Caesar cipher")
	flag.Parse()
}