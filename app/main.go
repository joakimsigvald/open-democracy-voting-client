// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"open-democracy-voting-client/lib/language"
	"open-democracy-voting-client/lib/url"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	t := language.Translator(r)
	u := url.UrlFactory(r)

	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<link rel='stylesheet' type='text/css' href='/styles/app.css'>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<h1 class='hdr'>"+t("headline")+"</h1>")
	fmt.Fprintln(w, "<h2>"+t("vote")+"</h2>")
	fmt.Fprint(w, "<a href='"+u("yes")+"' class='button'>"+t("yes")+"</a>")
	fmt.Fprint(w, "<a href='"+u("no")+"' class='button'>"+t("no")+"</a>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}
