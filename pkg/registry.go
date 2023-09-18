package pkg

func Vacuum(registry string, reserve int) {
	execGarbageCollect()
	cleanupOldImages(registry, reserve)
	execGarbageCollect()
}

func execGarbageCollect() {

}

func cleanupOldImages(registry string, reserve int) {
	images := getImages(registry)
	for i := range images {
		println(i)
		digests := getImageDigests(registry, images[i], reserve)
		println(len(digests))

	}
}

func getImages(registry string) (images []string) {
	return
}

func getImageDigests(registry string, image string, reserve int) (digest []ImageDigest) {

	return
}
