package request

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
)

func StopAll(req gorequest.Request, via []gorequest.Request) error {
	return http.ErrUseLastResponse
}
