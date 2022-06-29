package hdu_cas_helper

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
	"net/http"
	"strings"
)

type NewJW struct {
	cookie    map[string]string
	casStatus *LoginStatus
	err       error
}

func (s *NewJW) GetCookie() string {
	var str strings.Builder
	for _, value := range s.cookie {
		str.WriteString(value)
		str.WriteString("; ")
	}
	return str.String()
}

func (s *NewJW) Error() error {
	return s.err
}

func NewJWLogin(cas *LoginStatus) *NewJW {
	ret := &NewJW{
		cookie:    map[string]string{},
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
		Get(NEWJWLOGINURL).
		Set("User-Agent", UA).
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	route := getCookie(resp, "route")
	nextUrl := resp.Header.Get("Location") //跳转cas
	// 抓ticket
	resp, _, err = cli.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
		if req.URL.Path != "/jwglxt/xtgl/login_slogin.html" {
			return nil
		}
		return http.ErrUseLastResponse
	}).
		Get(nextUrl).
		Set("Cookie", cas.tgc).
		Set("User-Agent", UA).
		End()
	cas.expired = true
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	jsessionid := getCookie(resp, "JSESSIONID")

	nextUrl = resp.Header.Get("Location") //http://newjw.hdu.edu.cn/jwglxt/xtgl/login_slogin.html
	cli.
		Get(nextUrl).
		Set("User-Agent", UA).
		Set("Cookie", route+"; "+jsessionid).
		End()

	ret.cookie["JSESSIONID"] = jsessionid
	ret.cookie["route"] = route
	ret.casStatus.service = "newjw"
	return ret
}
