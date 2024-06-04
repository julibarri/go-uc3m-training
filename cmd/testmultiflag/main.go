package main

import (
	"flag"
	"fmt"
	"github.com/julibarri/go-uc3m-training/internal/helloworld"
	"github.com/julibarri/go-uc3m-training/internal/cliscanner"
	"github.com/julibarri/go-uc3m-training/internal/randrunes"
)

func init() {
	flag.Parse()
}

func main() {
	helloworld.HelloWorld()
	cliscanner.CliScanner()
	fmt.Println(randrunes.RandBytes(10, "abcdefghijklmn"))
	fmt.Println(randrunes.RandRunes(10, "\U0001f41b\U0001f42b\U0001f43b\U0001f44b\U0001f45b\U0001f46b\U0001f47b\U0001f48b\U0001f49b\U0001f4ab\U0001f4bb"))
}
