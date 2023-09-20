package pkg

import (
	"encoding/json"
	"log"
	"net/url"
	"sort"
	"time"
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

/*
List all images in a private registry-v2.

	https://docs.docker.com/registry/spec/api/#listing-repositories
	GET /v2/_catalog?n=<n from the request>&last=<last repository value from previous response>
	Args:
	  registry_url: The URL of the private registry.
	  headers = { "Authorization": "Basic " + urllib.parse.quote(username + ":" + password) }
	Returns:
	  A list of images in the registry. {"repositories":["cagalog1/cagalog2/image-name"]}
*/
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

/*
		GET /v2/<name>/tags/list?n=<n from the request>&last=<last tag value from previous response>
	    {"name":"catalog1/catalog2/image-name","tags":["prd-1","prd-2","prd-3"]}
*/
func getImageDigests(registry string, image string, reserve int) (digests []ImageDigest) {
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

		manifests := make(map[string][]string)

		for _, tag := range tags {
			log.Println(registry, image, tag)
			digest := getImageDigest(registry, image, tag)
			if tt, ok := manifests[digest.ManifestDigest]; ok {
				manifests[digest.ManifestDigest] = append(tt, tag)
				log.Println(">> Repeated tags:", image, manifests[digest.ManifestDigest])
				continue
			}
			manifests[digest.ManifestDigest] = []string{tag}
			log.Println(digest)
			digests = append(digests, digest)
		}
		log.Println(image, "authentic count", len(digests))
		sort.Slice(digests, func(i, j int) bool {
			return digests[i].Created.Compare(digests[j].Created) > 0
		})
		for _, v := range digests {
			log.Println(v.ToString())
		}
	}
	return
}

/*
		{registry}/v2/{image_name}/manifests/{image_tag}
	    json-body['config']['digest']
*/
func getImageDigest(registry string, image string, tag string) (digest ImageDigest) {
	url, _ := url.JoinPath(registry, "/v2/", image, "manifests", tag)
	rpHeader, rpBody, err := RequestRegistry(url, "GET")
	if err != nil {
		log.Panicln(err)
	} else {
		digest.Registry = registry
		digest.Image = image
		digest.Tag = tag
		digest.ManifestDigest = rpHeader.Get("Docker-Content-Digest")
		// d1.manifests_digest = json1['config']['digest']
		var jsdata map[string]interface{}
		err = json.Unmarshal(rpBody, &jsdata)
		if err != nil {
			log.Panicln(err)
		}
		// blobs digest
		//log.Println(">> config", jsdata["config"])
		config := ConvertInterfaceToDict(jsdata["config"])
		digest.BlobsDigest = config["digest"].(string)
		digest.Created = getDigestCreated(registry, image, digest.BlobsDigest)
	}
	return
}

/*
request url /v2/<name>/blobs/<digest>
*/
func getDigestCreated(registry string, image string, blobsDigest string) (created time.Time) {
	url, _ := url.JoinPath(registry, "/v2/", image, "blobs", blobsDigest)
	_, rpBody, err := RequestRegistry(url, "GET")
	if err != nil {
		log.Panicln(err)
	} else {
		// d1.manifests_digest = json1['config']['digest']
		var jsdata map[string]interface{}
		err = json.Unmarshal(rpBody, &jsdata)
		if err != nil {
			log.Panicln(err)
		}
		created1 := jsdata["created"]
		log.Println(">> created", created1)
		//config := ConvertInterfaceToDict(jsdata["config"])
		created, _ = time.Parse(time.RFC3339Nano, created1.(string))
	}
	return
}

/*
https://docs.docker.com/registry/spec/api/#deleting-an-image

	DELETE /v2/<name>/manifests/<reference>
*/
func DeleteImagManifest() {

}
