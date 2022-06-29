package hdu_cas_helper

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
)

func getCli(method, host, uri string, ssl bool) *gorequest.SuperAgent {
	var url string
	var proto string
	if ssl {
		url = fmt.Sprintf("https://%s%s", host, uri)
		proto = fmt.Sprintf("https://%s", host)
	} else {
		url = fmt.Sprintf("http://%s%s", host, uri)
		proto = fmt.Sprintf("http://%s", host)
	}
	UA := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

	switch method {
	case "get":
		{
			return gorequest.New().Get(url).
				RedirectPolicy(request.RedirectFunc).
				Set("User-Agent", UA).
				Set("Host", host).
				Set("Origin", proto).
				Set("Referer", url)
		}
	case "post":
		{
			return gorequest.New().Post(url).
				RedirectPolicy(request.RedirectFunc).
				Set("User-Agent", UA).
				Set("Host", host).
				Set("Origin", proto).
				Set("Referer", url)
		}
	}
	return nil
}
