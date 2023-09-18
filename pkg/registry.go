package pkg

import (
	"time"
)

type ImageDigest struct {
	Registry       string
	Image          string
	Tag            string
	ManifestDigest string
	BlobsDigest    string
	Created        time.Time
}

func Vacuum(registry string, reserve int) {
	execGarbageCollect()
	cleanupOldImages(registry, reserve)
	execGarbageCollect()
}

func execGarbageCollect() {

}

func cleanupOldImages(registry string, reserve int) {
	digests := getImages(registry, reserve)
	println(len(digests))
}

func getImages(registry string, reserve int) (digest []ImageDigest) {

	return
}
