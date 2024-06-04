package main

import (
	"flag"
	"github.com/julibarri/go-uc3m-training/internal/helloworld"
	"github.com/julibarri/go-uc3m-training/internal/cliscanner"
)

func init() {
	flag.Parse()
}

func main() {
	helloworld.HelloWorld()
	cliscanner.CliScanner()
}
