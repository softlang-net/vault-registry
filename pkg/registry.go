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
	/*
			GET /v2/<name>/tags/list?n=<n from the request>&last=<last tag value from previous response>
		    {"name":"zyb/prd/zyb-api-cms-hub","tags":["prd-1","prd-2","prd-3"]}
	*/
	url, _ := url.JoinPath(registry, "/v2/", image, "/tags/list")
	rpHeader, rpBody, err := RequestRegistry(url, "GET")
	if err != nil {
		log.Panicln(err)
	} else {
		for k := range rpHeader {
			log.Println(k, rpHeader.Get(k))
		}
		var jsdata map[string]interface{}
		err = json.Unmarshal(rpBody, &jsdata)
		if err != nil {
			log.Panicln(err)
		}
		log.Println(">> tags", jsdata["tags"])
		tags := ConvertInterfaceToStringSlice(jsdata["tags"])
		for _, v := range tags {
			log.Println(registry, image, v)
		}
	}
	return
}
