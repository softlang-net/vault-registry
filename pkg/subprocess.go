package pkg

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func RequestRegistry(url string, method string, authBasic string) (rpHeader http.Header, rpBody []byte, err error) {

	// Create a new HTTP request.
	DebugLog(method, url)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	if authBasic != "" {
		req.Header.Set("Authorization", "Basic "+authBasic)
	}

	// Make the request.
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		err1, _ := io.ReadAll(resp.Body)
		log.Panicln("http-code", resp.StatusCode, string(err1))
	}

	// Read the response body.
	rpBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
		return
	}

	rpHeader = resp.Header.Clone()
	return
}

func ShellCall(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			DebugLog(scanner.Text())
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

func ShellCallResult(name string, args ...string) (result string) {
	cmd := exec.Command(name, args...)
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)

	go func() {
		for scanner.Scan() {
			s1 := scanner.Text()
			DebugLog(s1)
			result += s1
		}
		done <- true
	}()
	cmd.Start()
	<-done
	err := cmd.Wait()
	if err != nil {
		DebugLog((err))
	}
	return
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
			DebugLog(scanner.Text())
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
