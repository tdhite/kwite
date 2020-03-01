/*
http_test.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package http

import (
	"testing"

	"github.com/tdhite/kwite/internal/tplfunc/funcs"
)

const (
	urlhost    string = "kwite-1.kwitenamespace"
	rewrite    string = "kwite.site"
	fullurl1   string = "kwite://kwite-1.kwitenamespace:8080/kwiteurl"
	rewritten1 string = "http://kwite.site:8080/kwiteurl"
	fullurl2   string = "kwite://kwite-1.kwitenamespace/kwiteurl"
	rewritten2 string = "http://kwite.site/kwiteurl"
)

func testRewrite(fullurl, rewritten string, t *testing.T) {
	u, err := GetRewrite(fullurl, "http")
	if err != nil {
		t.Log("Failed url rewrite!", err)
		t.Fail()
	}

	if u != rewritten {
		t.Logf("Failed url rewrite. Expected %s, but got %s.", rewritten, u)
		t.Fail()
	}
}

func testLoadRewriteRules(t *testing.T) {
	// NOTE: this depends on proper formation of the examples directory
	// in this project.
	p := "../../../examples/configs/rewrite"
	if err := initRewriteRules(p); err != nil {
		t.Log("Failed to initialize rewrite rules from "+p, err)
		t.Fail()
	}
}

func TestLoadFuncs(t *testing.T) {
	defer funcs.ClearTemplateFuncs()

	if err := LoadFuncs(); err != nil {
		t.Log("Failed to load functions!", err)
		t.Fail()
	}

	SetRewrite(urlhost, rewrite)
	testRewrite(fullurl1, rewritten1, t)
	testRewrite(fullurl2, rewritten2, t)

	testLoadRewriteRules(t)
}
