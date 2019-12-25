/*
funcs.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package funcs

import (
	html_template "html/template"
	"log"
	text_template "text/template"

	tplfuncs "github.com/tdhite/kwite/internal/tplfunc/funcs"
	tplhttp "github.com/tdhite/kwite/internal/tplfunc/http"
	tplmath "github.com/tdhite/kwite/internal/tplfunc/math"
	tplstring "github.com/tdhite/kwite/internal/tplfunc/string"
)

func TextTemplateFuncs() text_template.FuncMap {
	return text_template.FuncMap(tplfuncs.TemplateFuncs)
}

func HtmlTemplateFuncs() html_template.FuncMap {
	return html_template.FuncMap(tplfuncs.TemplateFuncs)
}

func loadFuncs() error {
	funcs := []func() error{
		tplhttp.LoadFuncs,
		tplmath.LoadFuncs,
		tplstring.LoadFuncs,
	}

	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	if err := loadFuncs(); err != nil {
		log.Println("Init error! All Template NOT loaded successfully")
	} else {
		log.Println("Initialized package funcs -- all Template functions loaded")
	}
}
