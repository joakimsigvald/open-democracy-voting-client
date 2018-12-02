package urlparser

import (
	"net/http"
	"strings"
)

func PathParameter(r *http.Request) string {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}
