package hdu_cas_helper

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type formReq struct {
	Rsa       string `form:"rsa"`
	Ul        int    `form:"ul"`
	Pl        int    `form:"pl"`
	Lt        string `form:"lt"`
	Execution string `form:"execution"`
	EventId   string `form:"_eventId"`
}

func (c *formReq) makeForm() string {
	var builder strings.Builder
	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)
	for k := 0; k < t.NumField(); k++ {
		builder.WriteString(fmt.Sprintf("%s=%v", t.Field(k).Tag.Get("form"), v.Field(k).Interface()))
		if k != t.NumField()-1 {
			builder.WriteString("&")
		}
	}
	return builder.String()
}

func CasPasswordLogin(uname, password string) *LoginStatus {
	ret := &LoginStatus{
		method:  METHODPASSWORD,
		tgc:     "",
		expired: false,
		service: "",
		err:     nil,
	}
	resp, body, _ := getCli("get", "cas.hdu.edu.cn", "/cas/login", true).End()
	regStr := `<input type="hidden" id="lt" name="lt" value="(?P<lt>.*?)" \/>[\s\S]*<input type="hidden" name="execution" value="(?P<execution>.*)" \/>[\s\S]*<input type="hidden" name="_eventId" value="(?P<_eventId>.*)" \/>`
	r := regexp.MustCompile(regStr)
	matches := r.FindStringSubmatch(body)
	if len(matches) != 4 {
		ret.err = errors.New("cas request err, maybe hdu service is done, please check")
		return ret
	}
	lt := matches[1]
	execution := matches[2]
	_eid := matches[3]
	req := formReq{
		Rsa:       strEnc(uname+password+lt, "1", "2", "3"),
		Ul:        len(uname),
		Pl:        len(password),
		Lt:        lt,
		Execution: execution,
		EventId:   _eid,
	}
	jsid := getCookie(resp, "JSESSIONID")

	if len(jsid) == 0 {
		ret.err = errors.New("get initial status failed, maybe hdu service is done, please check manually")
		return ret
	}

	resp, _, _ = getCli("post", "cas.hdu.edu.cn", "/cas/login", true).
		Set("Cookie", jsid).
		Type("form-data").
		Send(req.makeForm()).End()

	casTgc := getCookie(resp, "CASTGC")
	if len(casTgc) == 0 {
		ret.err = errors.New("cas login failed, maybe hdu service is done, please check manually")
		return ret
	}
	return &LoginStatus{
		method:  METHODPASSWORD,
		tgc:     casTgc,
		expired: false,
		service: "",
		err:     nil,
	}
}
