package test

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/softlang-net/vault-registry/pkg"
)

func TestCatalog(t *testing.T) {
	pkg.Vacuum(pkg.URL_REGISTRY, 10)
}

func TestConfig(t *testing.T) {
	t.Log("vault_registry", os.Getenv("vault_registry"))
}

func TestTemporary(t *testing.T) {
	//  dname, err := os.MkdirTemp("", "sampledir")
	if f, err := os.CreateTemp("", "sample"); err != nil {
		t.Error(err)
	} else {
		f.WriteString("hello, temporary file. To write log")
		t.Log("temporary file name", f.Name())
	}
}

func TestMap(t *testing.T) {
	digest := make(map[string]int)
	digest["hello"] = 1
	digest["hello"] = 2
	if v, ok := digest["hello"]; ok {
		t.Log("ok, hello =", v)
	}
}

func TestHttpGet(t *testing.T) {
	// Create a new HTTP request.
	req, err := http.NewRequest("GET", "http://localhost:9005/v2/_catalog", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Make the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Print the response body.
	log.Println(string(body))
}
