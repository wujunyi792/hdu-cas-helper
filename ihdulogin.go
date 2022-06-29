package hdu_cas_helper

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
	"strings"
)

type IHdu struct {
	cookie    map[string]string
	casStatus *LoginStatus
	err       error
}

func (s *IHdu) GetCookie() string {
	var str strings.Builder
	for _, value := range s.cookie {
		str.WriteString(value)
		str.WriteString("; ")
	}
	return str.String()
}

func (s *IHdu) Error() error {
	return s.err
}

func IHduLogin(cas *LoginStatus) *IHdu {
	ret := &IHdu{
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
	// 抓state
	resp, _, err := cli.RedirectPolicy(request.StopAll).
		Get(IHDULOGINURL).
		Set("User-Agent", UA).
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	nextUrl := resp.Header.Get("Location") //跳转cas
	// 抓ticket
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(nextUrl).
		Set("Cookie", cas.tgc).
		Set("User-Agent", UA).
		End()
	cas.expired = true
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	nextUrl = resp.Header.Get("Location") //跳ihdu, url携带ticket
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(nextUrl).
		Set("User-Agent", UA).
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}

	tpUp := getCookie(resp, "tp_up")
	if len(tpUp) == 0 {
		ret.err = errors.New("ihdu login failed, maybe hdu service is done, please check manually")
		return ret
	}
	// 其实已经有cookie了， 接下来操作为了更真实模拟浏览器跳转
	finalLoc := resp.Header.Get("Location")
	cli.Get(finalLoc).Set("Cookie", tpUp).End()
	ret.cookie["tp_up"] = tpUp
	ret.casStatus.service = "ihdu"
	return ret
}
