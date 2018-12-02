// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package language

import (
	"encoding/json"
	"io/ioutil"
)

var phrases map[string]map[string]string

const defaultLang = "sv"

func init() {
	phrases = loadPhrases()
}

func Translator(lang string) func(string) string {
	return func(term string) string {
		phrase := phrases[term]
		translation, ok := phrase[lang]
		if ok {
			return translation
		}
		translation, ok = phrase[defaultLang]
		if ok {
			return translation
		}
		return term
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
