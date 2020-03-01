/*
file.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package file

import (
	"io/ioutil"
	"log"
)

// Read the contents of a file and return the bytes therein.
func ReadFileBytes(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed reading file at %s: %s\n", path, err)
		return nil, err
	}
	return b, nil
}

// Read the contents of a file and return the contents as a string.
func ReadFileString(path string) (string, error) {
	if b, err := ReadFileBytes(path); err != nil {
		return "", err
	} else {
		return string(b), err
	}
}
