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
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/tdhite/kwite/internal/file"
)

type fileWatcher struct {
	watchingDir bool
	path        string
	watcher     *fsnotify.Watcher
}

var rulesWatcher *fileWatcher

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

func doWatches(dir, fname string, finished <-chan struct{}) {

	log.Println("http: starting watcher thread loop.")
	for {
		select {
		case event, ok := <-rulesWatcher.watcher.Events:
			if !ok {
				return
			}
			log.Println("http: rewrite rules file changed, trying to update kwite url mapping rules.")
			path := filepath.Join(dir, fname)
			if event.Op&fsnotify.Write == fsnotify.Write {
				if _, err := os.Stat(path); err == nil {
					if rulesWatcher.watchingDir {
						rulesWatcher.watchingDir = false
						log.Println("http: removing from watch: ", dir)
						rulesWatcher.watcher.Remove(dir)
						log.Println("http: adding to watch: ", path)
						rulesWatcher.watcher.Add(path)
					}
					if err := initRewriteRules(path); err != nil {
						log.Println(err)
					}
					log.Println("http: updated rewrite rules succeeded.")
				}
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				log.Println("http: rewrite rules file deleted, removing all rules.")
				setRewriteRules(nil)
				rulesWatcher.watchingDir = true
				log.Println("http: removing from watch: ", path)
				rulesWatcher.watcher.Remove(path)
				log.Println("http: adding to watch: ", dir)
				rulesWatcher.watcher.Add(dir)
			}
		case err, ok := <-rulesWatcher.watcher.Errors:
			if !ok {
				return
			}
			log.Println("http: ", err)
		case <-finished:
			log.Println("http: terminating file watcher.")
			rulesWatcher.watcher.Close()
			return
		}
	}
}

func WatchRewriteRules(path string) {
	// Load and setup rewrite map watcher
	initRewriteRules(path)

	if rulesWatcher != nil {
		rulesWatcher.watcher.Close()
	} else {
		rulesWatcher = &fileWatcher{
			path:        path,
			watchingDir: true,
		}
		// watch for future updates
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Println("http: ", err)
			return
		}
		rulesWatcher.watcher = watcher
	}

	dir, fname := filepath.Split(path)
	finished := make(chan struct{}, 1)
	go doWatches(dir, fname, finished)

	log.Println("http: adding file watch on ", dir)
	if err := rulesWatcher.watcher.Add(dir); err != nil {
		log.Println("http: ", err)
		return
	}

	// Try graceful shutdown (if possible)
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)
	go func() {
		// wait on signal to start the server shutdown
		log.Println("http: watcher waiting for os interrupt signal.")
		<-ossig
		log.Println("http: recieved signal to stop, shutting down file watcher.")

		// write (via closing) the finished channel to free up listeners
		close(finished)
	}()
}
