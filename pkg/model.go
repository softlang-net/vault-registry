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

type Catalog struct {
	Repositories []string `json:"repositories,omitempty"`
}
