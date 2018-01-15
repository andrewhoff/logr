package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/andrewhoff/logr"
	"github.com/andrewhoff/logr/config"
)

const hostname = "http://localhost:8080"

var (
	mode     = flag.String("mode", "read", "serve|read|write")
	priority = flag.Int("priority", 1, "priority level for log message possible values = [1,2,3]")
	msg      = flag.String("msg", "hello world", "message text for log")
)

func main() {
	flag.Parse()

	switch *mode {

	case "serve":
		logr.InitWithOpts(&config.Opts{Overwrite: true})
		logr.Serve()

	case "read":
		content, err := httpRead()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(content)

	case "write":
		if err := httpWrite(*priority, *msg); err != nil {
			log.Fatal(err)
		}

	case "severe":
		if err := httpWriteSevere(*msg); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("Usage: cli --mode=[serve|read|write] --priority=[1,2,3] --msg=\"hello world\"")

	}
}

func httpRead() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/read", hostname))
	if err != nil {
		return "", fmt.Errorf("Error encountered while trying to read: %v", err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func httpWrite(priority int, msg string) error {
	data := make(url.Values, 0)

	data.Add("priority", strconv.Itoa(priority))
	data.Add("msg", msg)

	resp, err := http.PostForm(fmt.Sprintf("%s/write", hostname), data)
	if err != nil {
		return fmt.Errorf("Error encountered while trying to write: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Write attempted with http status: %d\n", resp.StatusCode)
	return nil
}

func httpWriteSevere(msg string) error {
	data := make(url.Values, 0)

	data.Add("msg", msg)

	resp, err := http.PostForm(fmt.Sprintf("%s/write/severe", hostname), data)
	if err != nil {
		return fmt.Errorf("Error encountered while trying to write: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Write attempted with http status: %d\n", resp.StatusCode)
	return nil
}
