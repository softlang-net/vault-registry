package main

import "github.com/softlang-net/xRegistry/pkg"

func main() {
	println("hello xRegistry")
	url := "localhost:5000"

	pkg.Vacuum(url, 10)
}
