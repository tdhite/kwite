/*
rewrite.go

Copyright (c) 2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package http

import (
	"encoding/json"
	"errors"
	"log"
	neturl "net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/tdhite/kwite/internal/file"
)

type rewriteRules struct {
	mappings map[string]string
	lock     *sync.Mutex
}

var urlMap rewriteRules

func init() {
	urlMap.lock = &sync.Mutex{}
	setRewriteRules(make(map[string]string))
}

func setRewriteRules(m map[string]string) {
	urlMap.lock.Lock()
	urlMap.mappings = m
	urlMap.lock.Unlock()
}

// Load the rewrite rules from a JSON file
func initRewriteRules(path string) error {
	b, err := file.ReadFileBytes(path)
	if err != nil {
		log.Println("WARNING: failed to load rewrite rules from ", path)
		return err
	}

	m := make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Println("ERROR: failed to load rewrite rules.", err)
		return err
	}

	setRewriteRules(m)
	return nil
}

// Sets a rewrite (i.e., replacement) rule for the given path. If the with
// value is the empty string, the path replace mapping is deleted.
func SetRewrite(replace, with string) {
	urlMap.lock.Lock()
	if with == "" {
		delete(urlMap.mappings, replace)
	} else {
		urlMap.mappings[replace] = with
	}
	urlMap.lock.Unlock()
}

// Returns a rewritten url, using the urlMap table, to replace the 'hostname'
// part with the mapped entry withe the scheme changed to the supplied scheme.
// For example, the example url might  have the kwite-name.namespace portion
// replaced with the urlMap value for that as the map key and the scheme
// changed form kwite to the value of the scheme parameter:
// kwite://kwite-name.namespace/kwite-url.
func GetRewrite(url, scheme string) (string, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		log.Println("http: Invalid url: ", url, ".", err)
		return url, err
	}

	// if not a kwite scheme, no rewrite occurs
	if u.Scheme != "kwite" {
		return url, nil
	}

	r := urlMap.mappings[u.Hostname()]
	if r == "" {
		err := errors.New("http: Failed to find rewrite mapping for: " + u.Hostname())
		log.Println(err)
		return url, err
	}

	if u.Port() == "" {
		u.Host = r
	} else {
		u.Host = r + ":" + u.Port()
	}
	u.Scheme = scheme
	return u.String(), nil
}

func waitFileExists(path string) {
	// wait for the file to exist
	for {
		log.Println("Checking for file exists at " + path)
		_, err := os.Lstat(path)
		if os.IsNotExist(err) {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
}

func doWatch(watcher *fsnotify.Watcher, path string) {
	log.Println("http: clearing out any existing watches on ", path)

	// graceful shutdown term handling
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)

	for {
		waitFileExists(path)

		initRewriteRules(path)

		log.Println("http: adding file watch on ", path)
		if err := watcher.Add(path); err != nil {
			log.Println("http: ", err)
			break
		}

		log.Println("http: starting rewrite file watcher thread.")
		cont := true
		for cont {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					cont = false
				}
				log.Println("http: rewrite rules file changed, trying to update kwite url mapping rules.")
				if event.Op&fsnotify.Write == fsnotify.Write {
					if err := initRewriteRules(path); err != nil {
						log.Println(err)
					}
					log.Println("http: updated rewrite rules succeeded.")
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("http: rewrite rules file deleted, removing all rules and watches.")
					watcher.Remove(path)
					setRewriteRules(nil)
					cont = false
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					cont = false
				}
				log.Println("http: ", err)
			case <-ossig:
				log.Println("http: terminating rewrite file watcher.")
				watcher.Close()
				return
			}
		}
	}
}

func WatchRewriteRules(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("http: ", err)
		return
	}

	go doWatch(watcher, path)
}
