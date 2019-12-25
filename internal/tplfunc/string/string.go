/*
string.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package string

import (
	"strings"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

// Adds all methods from the strings package other than those related to
// strings.Builder, strings.Reader or those having a function as a parameter
// (e.g., FieldFunc).
func LoadFuncs() error {
	f := map[string]interface{}{
		"strCompare":        strings.Compare,
		"strContains":       strings.Contains,
		"strContainsAny":    strings.ContainsAny,
		"strContainsRune":   strings.ContainsRune,
		"strCount":          strings.Count,
		"strEqualFold":      strings.EqualFold,
		"strFields":         strings.Fields,
		"strHasPrefix":      strings.HasPrefix,
		"strHasSuffix":      strings.HasSuffix,
		"strIndex":          strings.Index,
		"strIndexAny":       strings.IndexAny,
		"strIndexByte":      strings.IndexByte,
		"strIndexRune":      strings.IndexRune,
		"strJoin":           strings.Join,
		"strLastIndex":      strings.LastIndex,
		"strLastIndexAny":   strings.LastIndexAny,
		"strLastIndexByte":  strings.LastIndexByte,
		"strRepeat":         strings.Repeat,
		"strReplace":        strings.Replace,
		"strReplaceAll":     strings.ReplaceAll,
		"strSplit":          strings.Split,
		"strSplitAfter":     strings.SplitAfter,
		"strSplitAfterN":    strings.SplitAfterN,
		"strSplitN":         strings.SplitN,
		"strTitle":          strings.Title,
		"strToLower":        strings.ToLower,
		"strToLowerSpecial": strings.ToLowerSpecial,
		"strToTitle":        strings.ToTitle,
		"strToTitleSpecial": strings.ToTitleSpecial,
		"strToUpper":        strings.ToUpper,
		"strToUpperSpecial": strings.ToUpperSpecial,
		"strToValidUTF8":    strings.ToValidUTF8,
		"strTrim":           strings.Trim,
		"strTrimLeft":       strings.TrimLeft,
		"strTrimPrefix":     strings.TrimPrefix,
		"strTrimRight":      strings.TrimRight,
		"strTrimSpace":      strings.TrimSpace,
		"strTrimSuffix":     strings.TrimSuffix,
		"strNewReplacer":    strings.NewReplacer,
	}

	return funcs.AddMethods(f)
}
