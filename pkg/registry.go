package pkg

import (
	"encoding/json"
	"log"
	"net/url"
)

func Vacuum(registry string, reserve int) {
	execGarbageCollect()
	cleanupOldImages(registry, reserve)
	execGarbageCollect()
}

func execGarbageCollect() {
	// bin/registry garbage-collect /etc/docker/registry/config.yml
	//ShellCall("registry", "garbage-collect", "/etc/docker/registry/config.yml")
}

func cleanupOldImages(registry string, reserve int) {
	catalog := getCatalog(registry)
	for _, img := range catalog.Repositories {
		log.Println(registry, img)
		digests := getImageDigests(registry, img, reserve)
		log.Println(len(digests))
	}
}

func getCatalog(registry string) (catalog Catalog) {
	url, _ := url.JoinPath(registry, "/v2/_catalog")
	rpHeader, rpBody, err := RequestRegistry(url, "GET")
	if err != nil {
		log.Panicln(err)
	} else {
		for k := range rpHeader {
			log.Println(k, rpHeader.Get(k))
		}
		err = json.Unmarshal(rpBody, &catalog)
		if err != nil {
			log.Panicln(err)
		}
	}
	return
}

func getImageDigests(registry string, image string, reserve int) (digest []ImageDigest) {

	return
}
