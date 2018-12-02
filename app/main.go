// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"open-democracy-voting-client/lib/language"
	"open-democracy-voting-client/lib/urlparser"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	lang := urlparser.PathParameter(r)
	t := language.Translator(lang)
	fmt.Fprintln(w, t("headline"))
}
