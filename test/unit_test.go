package test

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/softlang-net/vault-registry/pkg"
)

func TestYaml(t *testing.T) {
}

func TestCatalog(t *testing.T) {
	pkg.Vacuum(pkg.URL_REGISTRY, 10)
}

func TestConfig(t *testing.T) {
	t.Log("vault_registry", os.Getenv("vault_registry"))
	t1 := time.Now()
	t2 := t1.Add(-time.Hour * 48)
	t.Log(t1.Format(time.RFC3339), t2.Format(time.RFC3339), t1.Compare(t2))

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(aa)
	aa = aa[5:]
	t.Log(aa)
	aa = make([]int, 0)
	t.Log(aa)
}

func TestParseAddrPort(t *testing.T) {
	regex := regexp.MustCompile(`addr:\s*:\d+`)
	config := `
version: 1.0
http:
	addr: :5000
	header: aaa
health:
	storagedriver:
	  enabled: true
	  interval: 10s
	  threshold: 3
	`
	match := regex.FindString(config)
	t.Log("matched =", match)
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
