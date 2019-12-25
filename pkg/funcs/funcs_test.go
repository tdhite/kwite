/*
funcs_test.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package funcs

import (
	"testing"
)

func TestLoadFuncs(t *testing.T) {
	if err := loadFuncs(); err == nil {
		t.Log("Failed to load functions during init!", err)
		t.Fail()
	}
}
