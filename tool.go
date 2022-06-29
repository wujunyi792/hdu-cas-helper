package hdu_cas_helper

import (
	"github.com/parnurzeal/gorequest"
	"regexp"
	"strings"
)

func getCookie(resp gorequest.Response, name string) string {
	cookieStr := strings.Join(resp.Header.Values("Set-Cookie"), ";")
	regStr := `(?P<cookie>` + name + `=.*?);`
	r := regexp.MustCompile(regStr)
	matches := r.FindStringSubmatch(cookieStr)
	if len(matches) > 0 {
		return matches[len(matches)-1]
	}
	return ""
}
