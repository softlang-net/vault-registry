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

func ConvertInterfaceToStringSlice(interfaceValue interface{}) []string {
	if interfaceValue == nil {
		return nil
	}

	// If the interfaceValue is a slice of strings, return it directly.
	if stringSlice, ok := interfaceValue.([]string); ok {
		return stringSlice
	}

	// Otherwise, create a new slice of strings and iterate over the elements of the interfaceValue, converting each element to a string and adding it to the slice.
	var stringSlice []string
	for _, element := range interfaceValue.([]interface{}) {
		stringSlice = append(stringSlice, element.(string))
	}

	return stringSlice
}
