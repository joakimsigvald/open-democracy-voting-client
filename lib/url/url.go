package url

import (
	"net/http"
)

const defaultLang = "sv"

func GetLang(r *http.Request) string {
	lang := r.URL.Query().Get("lang")
	if len(lang) < 1 {
		return defaultLang
	}
	return lang
}

func UrlFactory(r *http.Request) func(string) string {
	lang := GetLang(r)
	query := "?lang=" + lang
	return func(path string) string {
		return "/" + path + query
	}
}
