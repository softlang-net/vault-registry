package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/softlang-net/vault-registry/pkg"
)

var (
	image   string
	reserve int
)

func init() {
	flag.StringVar(&image, "image", "", "-image=xxx:5000/abc/xyz:latest")
	flag.IntVar(&reserve, "reserve", 10, "-reserve=10")
}

func main() {
	// log.SetOutput(io.Discard)
	if isVacuumAnImage() {
		vacuumImage()
	} else {
		println("hello, welcome  vault-registry")
		url := pkg.URL_REGISTRY
		reserve := pkg.IMG_RESERVED
		pkg.Vacuum(url, reserve)
	}
}

func isVacuumAnImage() bool {
	//log.Panicln("there is no /bin/registry")
	println(strings.Join(os.Args, ","))
	for _, v := range os.Args {
		if v == "vacuum" {
			return true
		}
	}
	return false // do cleanup all
}

func vacuumImage() {
	flag.Parse()
	println(image)
	if !strings.Contains(image, "://") {
		image = "https://" + image
	}
	uri, err := url.Parse(image)
	if err != nil {
		log.Fatalln(err)
	} else {
		println(uri.Host, uri.Port(), uri.Path)
	}
	registry := uri.Scheme + uri.Host + ":" + uri.Port()

	pkg.VacuumAnImage(registry, uri.Path, reserve)
}
