/*
json.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package json

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

// Escapes JSON data via wrapper around https://godoc.org/encoding/json#HTMLEscape
func JsonHTMLEscape(data string) string {
	d := []byte(data)
	b := bytes.NewBufferString("")
	json.HTMLEscape(b, d)
	return b.String()
}

// Indents JSON via wrapper around https://godoc.org/encoding/json#Indent
func JsonIndent(data, prefix, indent string) (string, error) {
	d := []byte(data)
	b := bytes.NewBufferString("")
	if err := json.Indent(b, d, prefix, indent); err != nil {
		log.Println(err)
		return "", err
	}
	return b.String(), nil
}

// Unmartials JSON string into a Go type (e.g., map[string]interface{}) via
// wrapper around https://godoc.org/encoding/json#Unmarshal
func JsonToInterface(data string) (interface{}, error) {
	var j interface{}
	b := []byte(data)
	if err := json.Unmarshal(b, &j); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return j, nil
	}
}

// Unmartials JSON string into a Go type (e.g., map[string]interface{}) via
// wrapper around https://godoc.org/encoding/json#Marshal
func JsonToString(iface interface{}) (string, error) {
	if b, err := json.Marshal(iface); err != nil {
		log.Println(err)
		return "", err
	} else {
		return string(b), nil
	}
}

// Reports whether a string is valid JSON via wrapper around
// https://godoc.org/encoding/json#Valid
func JsonValid(j string) bool {
	return json.Valid([]byte(j))
}

// Adds all methods from the strings package other than those related to
// strings.Builder, strings.Reader or those having a function as a parameter
// (e.g., FieldFunc).
func LoadFuncs() error {
	f := map[string]interface{}{
		"jsonHTMLEscape":  JsonHTMLEscape,
		"jsonIndent":      JsonIndent,
		"jsonToInterface": JsonToInterface,
		"jsonToString":    JsonToString,
		"jsonValid":       JsonValid,
	}

	return funcs.AddMethods(f)
}
