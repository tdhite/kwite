/*
funcs_test.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package funcs

import (
	"testing"
)

func testFunc() string {
	return ""
}

func TestAddMethod(t *testing.T) {
	if err := AddMethod("test", testFunc); err != nil {
		t.Log("AddMathod failed to load a proper function!", err)
		t.Fail()
	}

	if err := AddMethod("test", testFunc); err == nil {
		t.Log("AddMathod allowed a duplicate function, which is prohibited!")
		t.Fail()
	}

	v := 1
	if err := AddMethod("v", v); err == nil {
		t.Log("AddMathod failed to prohibit loading a non-function!")
		t.Fail()
	}

	ClearTemplateFuncs()
}
