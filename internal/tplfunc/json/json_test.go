/*
json_test.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package json

import (
	"testing"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

func TestLoadFuncs(t *testing.T) {
	defer funcs.ClearTemplateFuncs()

	if err := LoadFuncs(); err != nil {
		t.Log("Failed to load functions!", err)
		t.Fail()
	}
}
