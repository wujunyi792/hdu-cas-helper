package hdu_cas_helper

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
	"regexp"
	"strings"
)

type HDQG struct {
	token         string
	hduhelpStatus *HduHelp
	err           error
}

func (s *HDQG) GetToken() string {
	return s.token
}

func (s *HDQG) Error() error {
	return s.err
}

func HDQGLogin(hh *HduHelp) *HDQG {
	ret := &HDQG{
		token:         "",
		hduhelpStatus: hh,
		err:           nil,
	}

	if hh.Error() != nil {
		ret.err = hh.Error()
		return ret
	}
	if hh.token == "" {
		ret.err = errors.New("token not found")
		return ret
	}

	cli := gorequest.New()
	resp, _, err := cli.RedirectPolicy(request.StopAll).
		Get(fmt.Sprintf(HDUHELPOAUTHLOGIN, HDQGAPPID, hh.GetToken())).
		Set("Upgrade-Insecure-Requests", "1").
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}

	nextUrl := resp.Header.Get("Location")
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(strings.Replace(nextUrl, "http:", "https:", 1)).
		Set("User-Agent", UA).
		End()
	//

	nextUrl = resp.Header.Get("Location")
	var token string
	if r := regexp.MustCompile("employment/(.*?)$").FindSubmatch([]byte(nextUrl)); len(r) > 0 {
		token = string(r[1])
	} else {
		ret.err = errors.New("can't get token")
		return ret
	}
	ret.token = token
	return ret
}
