/*
values.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package values

import (
	"fmt"
	"log"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

type variable struct {
	v interface{}
}

// Sets a the interface provided to the given value
func (v *variable) Set(value interface{}) string {
	log.Printf("values: Setting value of type %T to %#v", v.v, value)
	v.v = value
	return ""
}

// Sets a the interface provided to the given value
func (v *variable) Printf(format string) string {
	return fmt.Sprintf(format, v.v)
}

// Creates and returns a new variable
func new(initval interface{}) *variable {
	return &variable{v: initval}
}

// Adds all methods from the this package
func LoadFuncs() error {
	f := map[string]interface{}{
		"valuesNew": new,
	}

	return funcs.AddMethods(f)
}
