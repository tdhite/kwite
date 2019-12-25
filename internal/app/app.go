/*
app.go

Copyright (c) 2019 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tdhite/kwite/internal/globals"
)

const (
	configDirDefault          = "/configs"
	configDirUsage            = "where configuration (ConfigMap) is mounted"
	enableHttpFuncsDefault    = true
	enableHttpFuncsUsage      = "enable/disable http functions"
	enableMathFuncsDefault    = true
	enableMathFuncsUsage      = "enable/disable math functions"
	enableStringsFuncsDefault = true
	enableStringsFuncsUsage   = "enable/disable string functions"
	portDefault               = 8080
	portUsage                 = "port to bind to (default 8080)"
)

var configMap map[string]string

// Sets the string value from the environment key. If no
// key value exists, no action is taken.
func setEnvString(val *string, key string, dflt string) {
	str, ok := os.LookupEnv(key)
	if ok {
		*val = str
		fmt.Println("Set string val from environment key: ", key, "to ", str)
	}
}

// Sets the integer value from the environment key.
func setEnvInt(val *int, key string, dflt int) {
	str := ""
	sdflt := strconv.Itoa(dflt)
	setEnvString(&str, key, sdflt)
	if str != "" {
		if i, err := strconv.ParseInt(str, 0, 64); err != nil {
			fmt.Println("Conversion error to int for environment key: ", key)
			fmt.Println("Set to default value: ", dflt)
			*val = dflt
		} else {
			*val = int(i)
			fmt.Println("Set integer value from environment key ", key, "to value: ", i)
		}
	}
}

// Sets the boolean value from the environment key.
func setEnvBool(val *bool, key string, dflt bool) {
	str := ""
	var sdflt string
	if dflt {
		sdflt = "true"
	} else {
		sdflt = "false"
	}
	setEnvString(&str, key, sdflt)
	if str != "" {
		if b, err := strconv.ParseBool(str); err != nil {
			fmt.Println("Conversion error to bool for environment key: ", key)
			fmt.Println("Set to default value: ", dflt)
			*val = dflt
		} else {
			*val = b
			fmt.Println("Set integer value from environment key ", key, "to value: ", b)
		}
	}
}

func configureFromEnv() {
	log.Println("---- Setting Config From Environment ----")
	setEnvString(&globals.ConfigDir, "CONFIGDIR", configDirDefault)
	setEnvInt(&globals.Port, "PORT", portDefault)
}

// Initialize the flags processor with default values and help messages.
func initFlags() {
	log.Println("---- Setting Config From Command Line ----")
	flag.StringVar(&globals.ConfigDir, "configdir", configDirDefault, configDirUsage)
	flag.StringVar(&globals.ConfigDir, "c", configDirDefault, configDirUsage)
	flag.IntVar(&globals.Port, "port", portDefault, portUsage)
}

// Process application (command line) flags and environment. Environment takes
// precedence.
func Init() {
	initFlags()
	configureFromEnv()
	flag.Parse()
}

func init() {
	log.Println("Initialized app package.")
}
