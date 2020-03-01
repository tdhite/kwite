/*
file_test.go

Copyright (c) 2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package file

import (
	"encoding/json"
	"testing"
)

func TestLoadJSON(t *testing.T) {
	// NOTE: this depends on proper formation of the examples directory
	// in this project.
	p := "../../examples/configs/rewrite"
	if s, err := ReadFileBytes(p); err != nil {
		t.Log("Failed to load file "+p, err)
		t.Fail()
	} else if !json.Valid([]byte(s)) {
		t.Log("File loaded invalid JSON, expected valid.")
		t.Fail()
	}
}
