/*
config.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package app

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdhite/kwite/internal/globals"
)

// Read the contents of a (configmap) file and return the string value
func readFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed reading ConfigMap entry at %s: %s\n", path, err)
	}
	return string(b)
}

// Read the ConfigMap into a map from the mount point specified by the path
// Only certain keys of interest exist for this app: url, alive, ready and
// template. The url value is the url path to apply for 'handling' in the http
// server. The template is the go html or text template used for processing
// http requests.
func readConfigMap(path string) {
	configMap = make(map[string]string, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed opening config directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		path := filepath.Join(globals.ConfigDir, name)

		switch name {
		case "url":
			{
				configMap["url"] = strings.TrimSpace(readFile(path))
				log.Println("Key url: ", configMap["url"])
			}
		case "template":
			{
				configMap["template"] = path
				log.Println("Key template: ", configMap["template"])
			}
		case "ready":
			{
				configMap["ready"] = path
				log.Println("Ready template: ", configMap["ready"])
			}
		case "alive":
			{
				configMap["alive"] = path
				log.Println("Alive template: ", configMap["alive"])
			}
		default:
			log.Println("Ignoring ConfigMap key: ", name)
		}
	}
}
