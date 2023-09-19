package pkg

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func RequestRegistry(url string, method string) (rpHeader http.Header, rpBody []byte, err error) {

	// Create a new HTTP request.
	log.Println(method, url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	// Make the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
		return
	}

	defer resp.Body.Close()

	// Read the response body.
	rpBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	rpHeader = resp.Header.Clone()

	// Print the response body.
	log.Println(string(rpBody))
	return
}

func ShellCall(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		done <- true
	}()
	cmd.Start()
	<-done
	err := cmd.Wait()
	if err != nil {
		log.Panicln(err)
	}
}

func ShellCallResult(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		done <- true
	}()
	cmd.Start()
	<-done
	err := cmd.Wait()
	if err != nil {
		log.Println((err))
	}
	return ""
}

func ShellPipeStdin() {
	cmd := exec.Command("sh", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return
	}
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		done <- true
	}()

	go func() {
		defer stdin.Close()
		s1 := `
		echo 'hello world'
		pwd && hostname
		`
		io.WriteString(stdin, s1)
		io.WriteString(stdin, "echo 'done!'")
	}()
	err = cmd.Start()
	<-done
	if err != nil {
		log.Fatal(err)
	} else {
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	}
}
