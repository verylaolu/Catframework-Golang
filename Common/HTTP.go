package Common

import (
	"net/http"
)

func PostValue(r *http.Request,key string) string  {
	return r.FormValue(key)
}
func GetValue(r *http.Request,key string) string  {
	return r.URL.Query().Get(key)
}
func JsonValue(r *http.Request,key string) string  {
	return ""
}
