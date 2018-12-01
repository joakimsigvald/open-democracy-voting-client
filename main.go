// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	phrases := loadPhrases()
	fmt.Fprintln(w, phrases["headline"]["en"])
}

func loadPhrases() map[string]map[string]string {
	content, err := ioutil.ReadFile("resources/phrases.json")
	check(err)
	phrasesJson := string(content)
	phrases := map[string]map[string]string{}
	err = json.Unmarshal([]byte(phrasesJson), &phrases)
	check(err)
	return phrases
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
