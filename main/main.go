package main

import (
	"modtest"
	"modtest/util"
	"rsc.io/quote"
	"fmt"
)

func main() {
	util.Print("foobar")
	modtest.Hello()
	fmt.Println(quote.Hello())
}
