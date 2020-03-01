/*
realmain.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package app

import (
	"fmt"
	"log"
	"os"

	"github.com/tdhite/kwite/internal/globals"
	"github.com/tdhite/kwite/internal/server"
	"github.com/tdhite/kwite/internal/tplfunc/http"
)

var Template *server.Template

// Called by main, which is just a wrapper for this function. The reason
// is main can't directly pass back a return code to the OS.
func RealMain() int {
	Init()

	l := log.New(os.Stdout, "kwite: ", log.LstdFlags)

	// Load the configmap
	readConfigMap(globals.ConfigDir)

	// Initialize and watch for rewrite rules changes
	l.Println("Starting rewrite rules watcher on file ", configMap["rewrite"])
	http.WatchRewriteRules(configMap["rewrite"])

	Template, err := server.NewTemplate(configMap, server.IsFile, l)
	if err != nil {
		return 1
	}

	port := fmt.Sprintf(":%d", globals.Port)
	server := server.New(Template, configMap, port, l)
	server.StartServer(30)

	return 0
}
