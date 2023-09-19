package main

import (
	"github.com/softlang-net/vault-registry/pkg"
)

func main() {
	// log.SetOutput(io.Discard)
	checkCircumstance()
	println("hello, welcome  vault-registry")
	url := "http://127.0.0.1:5000"
	pkg.Vacuum(url, 10)
}

func checkCircumstance() {
	//log.Panicln("there is no /bin/registry")
}
