package main

import (
	"github.com/softlang-net/vault-registry/pkg"
)

func main() {
	// log.SetOutput(io.Discard)
	checkCircumstance()
	println("hello, welcome  vault-registry")
	url := pkg.URL_REGISTRY
	reserve := pkg.IMG_RESERVED
	pkg.Vacuum(url, reserve)
}

func checkCircumstance() {
	//log.Panicln("there is no /bin/registry")
}
