package hdu_cas_helper

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"time"
)

type QrLogin struct {
	login     *LoginStatus
	confirmed bool
	uuid      string
	url       string
}

func (q *QrLogin) GetScannerUrl() string {
	return q.url
}

func (q *QrLogin) GetUUID() string {
	return q.uuid
}

func (q *QrLogin) Login() *LoginStatus {
	type casRes struct {
		RedirectUrl string `json:"redirect_url"`
	}
	var res casRes
	resp, _, err := gorequest.New().
		Get("https://cas.hdu.edu.cn/cas/checkQRCodeScan?uuid=" + q.uuid).
		EndStruct(&res)
	if err != nil {
		q.login.err = err[0]
		return q.login
	}
	casTgc := getCookie(resp, "CASTGC")
	if len(casTgc) == 0 {
		q.login.err = errors.New("login failed, maybe hdu service is done, please check manually")
		return q.login
	}
	q.login.tgc = casTgc
	return q.login
}

func (q *QrLogin) AsyncLogin(interval time.Duration, retry int) *LoginStatus {
	for {
		s := q.Login()
		retry--
		if s.Error() == nil || retry <= 0 {
			return s
		}
		time.Sleep(interval)
	}
}

func QrCodeLogin() *QrLogin {
	uid := uuid.NewV4().String()
	return &QrLogin{
		login: &LoginStatus{
			method:  METHODQRCODE,
			tgc:     "",
			expired: false,
			service: "",
			err:     nil,
		},
		confirmed: false,
		uuid:      uid,
		url:       "https://cas.hdu.edu.cn/cas/qyQrLogin?uuid=" + uid,
	}
}
