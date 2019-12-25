/*
kwite.go

Copyright (c) 2019 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package main

import (
	"os"

	"github.com/tdhite/kwite/internal/app"
)

func main() {
	// Delegate to realMain so defered operations can happen (os.Exit exits
	// the program without servicing defer statements)
	os.Exit(app.RealMain())
}
