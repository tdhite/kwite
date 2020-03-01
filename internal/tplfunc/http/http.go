/*
http.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package http

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	nethttp "net/http"
	"strings"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

func doRequest(op, url, body string, headers []string) (string, error) {
	// before even starting, assure headers has even number of members
	if len(headers)%2 != 0 {
		err := errors.New("Header istrings has odd number of members; must be even")
		log.Println(err)
		return "", nil
	}

	// TODO: once https is supported, fix this via parameter passed scheme
	realUrl, err := GetRewrite(url, "http")
	if err != nil {
		log.Println(err)
		return "", err
	}

	c := nethttp.Client{}

	var b io.ReadCloser = nil
	if body != "" {
		b = ioutil.NopCloser(strings.NewReader(body))
	}

	req, err := nethttp.NewRequest(op, realUrl, b)
	if err != nil {
		log.Println(err)
		return "", err
	}

	for i := 0; i < len(headers); i += 2 {
		log.Println("http:  setting header as: " + headers[i] + ": " + headers[i+1])
		req.Header.Set(headers[i], headers[i+1])
	}

	log.Printf("http: performing %s on %s with headers: %v\n", op, realUrl, req.Header)
	rsp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rsp.Body.Close()

	// pull out the body and return it or error
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Print("Failed reading response body for url: " + realUrl)
		log.Println(err)
		return "", err
	}

	return string(rspBody), nil
}

func Get(url, body string, headers ...string) (string, error) {
	return doRequest("GET", url, body, headers)
}

func Delete(url, body string, headers ...string) (string, error) {
	return doRequest("DELETE", url, body, headers)
}

func Patch(url, body string, headers ...string) (string, error) {
	return doRequest("PATCH", url, body, headers)
}

func Post(url, body string, headers ...string) (string, error) {
	return doRequest("POST", url, body, headers)
}

// Adds all methods from this package to the template functions.
func LoadFuncs() error {
	f := map[string]interface{}{
		"httpDelete": Delete,
		"httpGet":    Get,
		"httpPatch":  Patch,
		"httpPost":   Post,
	}

	return funcs.AddMethods(f)
}
