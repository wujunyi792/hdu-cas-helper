package request

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
)

func RedirectFunc(req gorequest.Request, via []gorequest.Request) error {
	return http.ErrUseLastResponse
}
