package test

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/softlang-net/vault-registry/pkg"
)

func TestCatalog(t *testing.T) {
	pkg.Vacuum("http://localhost:5000", 10)
}

func TestHttpGet(t *testing.T) {
	// Create a new HTTP request.
	req, err := http.NewRequest("GET", "http://localhost:5000/v2/_catalog", nil)
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
