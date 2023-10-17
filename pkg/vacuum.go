package pkg

func VacuumAnImage(registry, image string, reserve int) {
	digests := getImageDigests(registry, image, reserve)
	deleteImagManifest(digests)
}
