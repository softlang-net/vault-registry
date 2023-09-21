package pkg

import (
	"encoding/base64"
	"flag"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

func DebugLog(v ...any) {
	log.Println(v...)
}

var (
	// Initialize a constant from the `PORT` environment variable.
	URL_REGISTRY string = getConfigString("registry", "vault_registry", "http://localhost:5000")
	IMG_RESERVED int    = getConfigInt("reserved", "vault_reserved", 500)
)

/*
get config value from flag by key, or get from os.environment by keyOfEnv
*/
func getConfigString(key, keyOfEnv, valDefault string) (value string) {
	value = *flag.String(key, "", key)
	if value == "" {
		value = os.Getenv(keyOfEnv)
		if value == "" {
			value = valDefault
		}
	}
	return
}

func getConfigInt(key, keyOfEnv string, valDefault int) (value int) {
	s1 := *flag.String(key, "", key)
	if s1 == "" {
		s1 = os.Getenv(keyOfEnv)
		if s1 == "" {
			value = valDefault
			return
		}
	}
	value, err := strconv.Atoi(s1)
	if err != nil {
		log.Panicln(err)
	}
	return
}

type ImageDigest struct {
	Registry       string
	Image          string
	Tag            string
	ManifestDigest string
	BlobsDigest    string
	Created        time.Time
}

func (image *ImageDigest) ToString() string {
	s, _ := url.JoinPath(image.Registry, image.Image, image.Tag, image.Created.Format(time.RFC3339))
	return s
}

type Catalog struct {
	Repositories []string `json:"repositories,omitempty"`
}

func Base64EncodeAuthentication(username, password string) string {
	credentials := username + ":" + password
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return encodedCredentials
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

func ConvertInterfaceToDict(interfaceValue interface{}) map[string]interface{} {
	if interfaceValue == nil {
		return nil
	}

	// If the interfaceValue is a slice of strings, return it directly.
	if stringSlice, ok := interfaceValue.(map[string]interface{}); ok {
		return stringSlice
	}
	return nil
}
