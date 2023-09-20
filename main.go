package main

import (
	"github.com/softlang-net/vault-registry/pkg"
)

func main() {
	// log.SetOutput(io.Discard)
	checkCircumstance()
	println("hello, welcome  vault-registry")
	url := pkg.URL_REGISTRY
	pkg.Vacuum(url, 10)
}

func checkCircumstance() {
	//log.Panicln("there is no /bin/registry")
}
