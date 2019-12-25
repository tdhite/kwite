/*
server_test.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

var (
	// Tests will bind servers to this and the next follow port
	testPort = 10080
)

const (
	testTemplate      = "x is {{ .x }}"
	testAliveAndReady = "OK!"
)

// Runs the http server and client test:
//    tmpl is the Template object
//    m is the map containing template information
//    p is a secondary url path to append (e.g., 'ready' or 'alive'
//    l is the logger to use for any log output
func test(t *testing.T, tmpl Template, m map[string]string, port int, p string, l *log.Logger) {
	// build a client
	url := "http://" + path.Join(fmt.Sprintf("localhost:%d", port), m["url"], p)
	t.Log("Querying at: " + url)
	client := &http.Client{}

	// start the server
	portSpec := fmt.Sprintf(":%d", port)
	t.Log("Starting server at: " + portSpec)
	server := New(tmpl, m, portSpec, l)
	go server.StartServer(5)

	// loop variables
	var resp *http.Response
	var err error
	j := []byte(`{"x": 12}`)

	// cheap retry loop to cover generally uninteresting http errors.
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", url, bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "text/plain")

		resp, err = client.Do(req)
		if err != nil {
			if strings.HasSuffix(err.Error(), ": connection refused") {
				// server isn't ready, try again
				t.Log("Failed connection (will retry): ", err)
			} else {
				t.Log("Failed to perform http GET (will retry): ", err)
			}
		}

		// delay for a tenth of a second before retry
		time.Sleep(time.Millisecond * 100)
	}

	if err != nil {
		t.Log("Failed to perform http GET. ", err)
		t.Fail()
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if p != "" {
		if string(body) != testAliveAndReady {
			t.Log("Test failed: got incorrect alive/ready response: ", string(body))
			t.Fail()
		}
	} else if string(body) != "x is 12" {
		t.Log("Test failed: expected 'x is 12' but got ", string(body))
		t.Fail()
	}
}

func writeTestFile(t *testing.T, name string, tmpl []byte) (string, error) {
	f, err := ioutil.TempFile("", name)
	if err != nil {
		t.Log(err)
		t.Fail()
		return "", err
	}

	if _, err := f.Write(tmpl); err != nil {
		t.Log(err)
		t.Fail()
		defer os.Remove(f.Name())
		return "", err
	}

	if err := f.Close(); err != nil {
		t.Log(err)
		t.Fail()
		defer os.Remove(f.Name())
		return "", err
	}

	return name, nil
}

func getConfigMap() map[string]string {
	m := make(map[string]string)
	m["url"] = "/kwite"
	m["template"] = string(testTemplate)
	m["ready"] = string(testAliveAndReady)
	m["alive"] = string(testAliveAndReady)
	return m
}

func TestFile(t *testing.T) {
	l := log.New(os.Stdout, "kwite-test: ", log.LstdFlags)
	m := getConfigMap()

	if name, err := writeTestFile(t, "kwite", []byte(string(testTemplate))); err != nil {
		return
	} else {
		defer os.Remove(name)
	}

	if name, err := writeTestFile(t, "ready", []byte(string(testTemplate))); err == nil {
		return
	} else {
		defer os.Remove(name)
	}

	if name, err := writeTestFile(t, "alive", []byte(string(testTemplate))); err == nil {
		return
	} else {
		defer os.Remove(name)
	}

	tmpl, err := NewTemplate(m, IsFile, l)
	if err != nil {
		t.Log("Template creation from file failed!", err)
		t.Fail()
	}

	test(t, tmpl, m, testPort+3, "", l)
	test(t, tmpl, m, testPort+4, "kwiteready", l)
	test(t, tmpl, m, testPort+5, "kwitealive", l)
}

func TestString(t *testing.T) {
	l := log.New(os.Stdout, "kwite-test: ", log.LstdFlags)
	m := getConfigMap()

	tmpl, err := NewTemplate(m, IsString, l)
	if err != nil {
		t.Log("Template creation failed!", err)
		t.Fail()
	}

	test(t, tmpl, m, testPort, "", l)
	test(t, tmpl, m, testPort+1, "kwiteready", l)
	test(t, tmpl, m, testPort+2, "kwitealive", l)
}
