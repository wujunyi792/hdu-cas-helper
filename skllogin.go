package hdu_cas_helper

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/wujunyi792/hdu-cas-helper/util/request"
	"regexp"
)

type Skl struct {
	token     string
	casStatus *LoginStatus
	err       error
}

func (s *Skl) GetToken() string {
	return s.token
}

func (s *Skl) Error() error {
	return s.err
}

func SklLogin(cas *LoginStatus) *Skl {
	ret := &Skl{
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
	// 抓state
	resp, _, err := cli.RedirectPolicy(request.StopAll).
		Get(SKLLOGINURL + "/api/cas/login?state=&index=&ticket=").
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	casUrl := resp.Header.Get("Location")
	// 抓ticket
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(casUrl).
		Set("Cookie", cas.tgc).
		End()
	cas.expired = true
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	sklUrl := resp.Header.Get("Location")
	resp, _, err = cli.RedirectPolicy(request.StopAll).
		Get(sklUrl).
		End()
	if len(err) != 0 {
		ret.err = err[0]
		return ret
	}
	finalLoc := resp.Header.Get("Location")
	regStr := `token=(.{36}?)[&]*.*?$`
	r := regexp.MustCompile(regStr)
	matches := r.FindStringSubmatch(finalLoc)
	if len(matches) != 2 {
		ret.err = errors.New("skl login failed, maybe skl service is done, please check manually")
		return ret
	}
	ret.token = matches[len(matches)-1]
	ret.casStatus.service = "skl"
	return ret
}
