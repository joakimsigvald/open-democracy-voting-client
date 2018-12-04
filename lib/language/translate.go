// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package language

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"open-democracy-voting-client/lib/url"
)

var phrases map[string]map[string]string

func init() {
	phrases = loadPhrases()
}

func Translator(r *http.Request) func(string) string {
	lang := url.GetLang(r)
	return func(term string) string {
		phrase := phrases[term]
		translation, ok := phrase[lang]
		if ok {
			return translation
		}
		return term + ":" + lang
	}
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
