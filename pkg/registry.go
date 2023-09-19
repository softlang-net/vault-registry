package pkg

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
	images := getImages(registry)
	for _, img := range images {
		println(img)
		digests := getImageDigests(registry, img, reserve)
		println(len(digests))
	}
}

func getImages(registry string) (images []string) {

	return
}

func getImageDigests(registry string, image string, reserve int) (digest []ImageDigest) {

	return
}
