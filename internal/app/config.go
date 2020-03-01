/*
config.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdhite/kwite/internal/file"
	"github.com/tdhite/kwite/internal/globals"
)

// Read the ConfigMap into a map from the mount point specified by the path
// Only certain keys of interest exist for this app: url, alive, ready and
// template. The url value is the url path to apply for 'handling' in the http
// server. The template is the go html or text template used for processing
// http requests.
func readConfigMap(path string) {
	configMap = make(map[string]string, 0)

	// setup default rewrite rules file
	configMap["rewrite"] = filepath.Join(globals.ConfigDir, "rewrite")

	f, err := os.Open(path)
	if err != nil {
		log.Printf("Failed opening config directory: %s", err)
		return
	}
	defer f.Close()

	list, _ := f.Readdirnames(0)
	for _, name := range list {
		path := filepath.Join(globals.ConfigDir, name)

		switch name {
		case "url":
			{
				if s, err := file.ReadFileString(path); err != nil {
					log.Println("Defaulting key url to kwite because of failed read of "+path, err)
					configMap["url"] = "/kwite"
				} else {
					configMap["url"] = strings.TrimSpace(s)
				}
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
		case "rewrite":
			{
				configMap["rewrite"] = path
				log.Println("Rewrite template: ", configMap["rewrite"])
			}
		default:
			log.Println("Ignoring ConfigMap key: ", name)
		}
	}
}
