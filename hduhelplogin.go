package hdu_cas_helper

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
	"regexp"
)

type HduHelp struct {
	token     string
	casStatus *LoginStatus
	err       error
}

func (s *HduHelp) GetToken() string {
	return s.token
}

func (s *HduHelp) Error() error {
	return s.err
}

func HduHelpLogin(cas *LoginStatus) *HduHelp {
	ret := &HduHelp{
		token:     "",
		casStatus: cas,
		err:       nil,
	}

	if cas.Error() != nil {
		ret.err = cas.Error()
		return ret
	}
	if cas.expired {
		ret.err = errors.New("cas ticket has been used or expired, please login cas again")
		return ret
	}

	cli := gorequest.New()
	resp, _, err := cli.RedirectPolicy(request.StopAll).
		Get(HDUHELPLOGINURL).
		Param("clientID", "app").
		Param("school", "hdu").
		Param("redirect", HDUHELPAPPURL).
		Set("User-Agent", UA).
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}

	nextUrl := resp.Header.Get("Location")
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(nextUrl).
		Set("Cookie", cas.tgc).
		Set("User-Agent", UA).
		End()
	//

	nextUrl = resp.Header.Get("Location")
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(nextUrl).
		Set("Cookie", cas.tgc).
		Set("User-Agent", UA).
		End()
	//
	nextUrl = resp.Header.Get("Location")
	var token string
	if r := regexp.MustCompile("auth=(.*?)$").FindSubmatch([]byte(nextUrl)); len(r) > 0 {
		token = string(r[1])
	} else {
		ret.err = errors.New("can't get token")
		return ret
	}
	cas.expired = true
	ret.casStatus.service = "hduhelp"
	ret.token = token
	return ret
}
