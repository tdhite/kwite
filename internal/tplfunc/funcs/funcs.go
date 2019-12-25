/*
funcs.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package funcs

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var TemplateFuncs map[string]interface{}

func init() {
	ClearTemplateFuncs()
}

// Clears the function table.
func ClearTemplateFuncs() {
	TemplateFuncs = make(map[string]interface{})
}

// Adds a method to the template functions.
// Returns an error if the interface is not a bona fide method or
// a method already exists with the provided name.
func AddMethod(name string, method interface{}) error {
	kind := reflect.TypeOf(method).Kind()
	if kind.String() != "func" {
		msg := fmt.Sprintf("Attempt to add a non-method (%s) to the function table.", kind.String())
		err := errors.New(msg)
		log.Println(err)
		return err
	}

	if _, ok := TemplateFuncs[name]; ok {
		msg := fmt.Sprintf("Attempt to add a non-method (%s) that already exists.", name)
		err := errors.New(msg)
		log.Println(err)
		return err
	}

	TemplateFuncs[name] = method
	return nil
}

// Adds all method in the map to the template functions.
// The function returns on first encounter of any error.
func AddMethods(funcs map[string]interface{}) error {
	for k, v := range funcs {
		if err := AddMethod(k, v); err != nil {
			return err
		}
	}

	return nil
}
